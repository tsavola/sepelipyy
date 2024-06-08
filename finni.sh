#!/bin/sh

curl https://kaino.kotus.fi/lataa/nykysuomensanalista2024.csv | cut -f1 | grep -E '^[abcdefghijklmnopqrstuvwxyzåäö]+$' | grep -E '.{4}' | sort -u > words.txt
