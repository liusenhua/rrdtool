rrdtool examples
================

Some typical case examples.

A single *install* script do below things automatically:

    - Create rrdtool database to save data.

    - Create a cron job to update data in rrdtool database.

        - The cron job use gawk to extract data.

    - Create another cron job to draw graph periodicity

The *uninstall* script clean everything.

## Dependency
    - RRDTool

        `sudo apt-get install rrdtool` for ubuntu, For other os, refer http://oss.oetiker.ch/rrdtool/index.en.html

    - gawk

    - Go environment

## Installation

    For each script, set the proper paths of the variables at the top of the script
    
        $(gawk)
        
        $(rrdtool)

    For each example(take 'ping' as example)
    
        cd ping
    
        chmod 755 ./install
    
        ./install
