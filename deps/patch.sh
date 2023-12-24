#!/bin/bash
file_path="./v8/third_party/icu/BUILD.gn"
replace_str="\"standard__\""
sed -i "s/\"standard\"/$replace_str/" $file_path
rm -rf "./v8/third_party/icu/.git/"
echo "Patch icu success"