placeholder
===========

A simple image placeholder service like [placehold.it](http://placehold.it).


# Usage

## Basic

    go run placeholder.go

It will run placeholder on port 8080, then you can use it like this:

    http://localhost:8080/350x150

## Size

    width x height


Height is optional, if no height is specified the image will be a square.

## Color

    http://localhost:8080/250/ffffff/000000

## Text

    http://localhost:8080/350x150?text=Golang+Placeholder


# TODO

+ Text break line support