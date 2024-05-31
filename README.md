# fluentbit_pipeline
End-to-end FluentBit Pipeline written in GoLang. Demonstrates utilizing FluentBit to ingest data from a dynamically-updating source using only configuration files.

Include all three files in whatever workspace/directory you are working in. Make sure to edit the configuration file to include your exact path. Fluent bit must be installed for this to work. Cd into the directory and run these commands in different terminals: 

go run source.go
tail -f logfile.log
fluent-bit -c expipeline.conf

After running the first two commands, you should see the log entries being appended with the current time each second (in the terminal where the tail command was run). This shows the progress of the logfile and confirms that the log entry script is working properly (source.go). After the third command you should see the log entries in stdout through fluent bit which confirms that the fluent bit pipeline is working properly.
