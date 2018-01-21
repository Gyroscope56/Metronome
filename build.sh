#!/bin/bash

go build
tar cvf output.tar ./index.html ./favicon.ico audio_stuff/ Metronome

