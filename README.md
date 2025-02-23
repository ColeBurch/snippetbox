# snippetbox

# setup Tailwind
1. Clone the repository
2. Setup tailwind css
    1. run the following command to install the tailwindcss cli replacing the end of the command with the correct file name for your os. "curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/<the_file_name_for_your_os>"
    2. rename the installed file to tailwindcss for convenvience running the following command "mv <the_file_name_for_your_os> tailwindcss"
    3. run the following command to make the file executable "chmod +x tailwindcss"
    4. run the following command in a seperate terminal to start the tailwindcss cli watcher "./tailwindcss -i ./ui/static/css/input.css -o ./ui/static/css/output.css --watch"
