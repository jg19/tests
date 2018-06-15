#!venv/bin/python3
import time
import nlpnet
tagger = nlpnet.POSTagger('pos-pt/', language='pt')

while True:
    print("Enter some text:")
    text = input()
    resp = tagger.tag(text)
    print(resp)
    