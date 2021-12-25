#!/usr/bin/env python3
from pathlib import Path
from collections import defaultdict


strip_chars = ('"', "path:", ',')


if __name__ == '__main__':
    content = (Path(__file__).parent / Path('./paths.in')).read_text()
    for line in content.splitlines():
        if "path:" in line:
            for to_strip in strip_chars:
                line = line.replace(to_strip, '')
            print(line)
