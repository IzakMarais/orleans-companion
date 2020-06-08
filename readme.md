# Orléans helper site

This is a fan made companion site for the [Orléans](https://boardgamegeek.com/boardgame/164928/orleans) boardgame.

## Setup helper

The setup in Orléans normally requires

- different types of goods tiles to be _shuffled together_,
- certain quantities randomly selected based on player counts
- certain quantities randomly removed based on player counts
- and the remaining tiles _sorted by type again_ to form a supply.

The setup helper aims to bypass this shuffling and resorting by calculating randomized setup numbers for selected and removed goods tiles.

## Use

The site is hosted at [orleans-companion.appspot.com](http://orleans-companion.appspot.com) on Google App engine.

## Development

### Google App engine usage reminder

- Download and install [gcloud sdk](https://cloud.google.com/sdk/docs/)
- Authenticate my account (may happen as part of installation):

      gcloud init

- As per the [quick start](https://cloud.google.com/appengine/docs/standard/go/quickstart), after installing app-engine-go you can run a local development server:

      gcloud components install app-engine-go
      dev_appserver.py app.yaml
      # If the above doesn't work try
      python <absolute path to dev_appserver.py> app.yaml

- Deploy using

      gcloud app deploy --project orleans-companion
