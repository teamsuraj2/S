package platforms

import (
        "context"
        "errors"
        "fmt"
        "io"
        "net/http"
        "os"

        "github.com/amarnathcjd/gogram/telegram"
                        "github.com/TheTeamVivek/YukkiMusic/internal/state"
)

var (
        apiBase = "https://youtubify.me"
        apiKey  = os.Getenv("YT_API_KEY")
)
const PlatformYoutubify state.PlatformName = "xyz"
type YoutubifyPlatform struct{}

func init() {
        addPlatform(100, PlatformYoutubify, &YoutubifyPlatform{})
}

func (*YoutubifyPlatform) Name() state.PlatformName {
        return PlatformYoutubify
}

func (*YoutubifyPlatform) IsValid(query string) bool {
        return false
}

func (*YoutubifyPlatform) GetTracks(query string) ([]*state.Track, error) {
        return nil, errors.New("Youtubify is a direct download platform")
}
