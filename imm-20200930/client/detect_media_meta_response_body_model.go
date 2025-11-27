// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMediaMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*Address) *DetectMediaMetaResponseBody
	GetAddresses() []*Address
	SetAlbum(v string) *DetectMediaMetaResponseBody
	GetAlbum() *string
	SetAlbumArtist(v string) *DetectMediaMetaResponseBody
	GetAlbumArtist() *string
	SetArtist(v string) *DetectMediaMetaResponseBody
	GetArtist() *string
	SetAudioStreams(v []*AudioStream) *DetectMediaMetaResponseBody
	GetAudioStreams() []*AudioStream
	SetBitrate(v int64) *DetectMediaMetaResponseBody
	GetBitrate() *int64
	SetComposer(v string) *DetectMediaMetaResponseBody
	GetComposer() *string
	SetDuration(v float64) *DetectMediaMetaResponseBody
	GetDuration() *float64
	SetFormatLongName(v string) *DetectMediaMetaResponseBody
	GetFormatLongName() *string
	SetFormatName(v string) *DetectMediaMetaResponseBody
	GetFormatName() *string
	SetLanguage(v string) *DetectMediaMetaResponseBody
	GetLanguage() *string
	SetLatLong(v string) *DetectMediaMetaResponseBody
	GetLatLong() *string
	SetPerformer(v string) *DetectMediaMetaResponseBody
	GetPerformer() *string
	SetProduceTime(v string) *DetectMediaMetaResponseBody
	GetProduceTime() *string
	SetProgramCount(v int64) *DetectMediaMetaResponseBody
	GetProgramCount() *int64
	SetRequestId(v string) *DetectMediaMetaResponseBody
	GetRequestId() *string
	SetSize(v int64) *DetectMediaMetaResponseBody
	GetSize() *int64
	SetStartTime(v float64) *DetectMediaMetaResponseBody
	GetStartTime() *float64
	SetStreamCount(v int64) *DetectMediaMetaResponseBody
	GetStreamCount() *int64
	SetSubtitles(v []*SubtitleStream) *DetectMediaMetaResponseBody
	GetSubtitles() []*SubtitleStream
	SetTitle(v string) *DetectMediaMetaResponseBody
	GetTitle() *string
	SetVideoHeight(v int64) *DetectMediaMetaResponseBody
	GetVideoHeight() *int64
	SetVideoStreams(v []*VideoStream) *DetectMediaMetaResponseBody
	GetVideoStreams() []*VideoStream
	SetVideoWidth(v int64) *DetectMediaMetaResponseBody
	GetVideoWidth() *int64
}

type DetectMediaMetaResponseBody struct {
	// The addresses.
	//
	// This parameter is returned only when address information is detected.
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The album.
	//
	// example:
	//
	// unable
	Album *string `json:"Album,omitempty" xml:"Album,omitempty"`
	// The album artist.
	//
	// example:
	//
	// unable
	AlbumArtist *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	// The artist.
	//
	// example:
	//
	// unable
	Artist *string `json:"Artist,omitempty" xml:"Artist,omitempty"`
	// The audio streams.
	AudioStreams []*AudioStream `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	// The bitrate. Unit: bit/s.
	//
	// example:
	//
	// 13164131
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The composer.
	//
	// example:
	//
	// unable
	Composer *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	// The total duration of the video. Unit: seconds.
	//
	// example:
	//
	// 15.263000
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The full format name.
	//
	// example:
	//
	// QuickTime / MOV
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	// The abbreviated format name.
	//
	// example:
	//
	// mov,mp4,m4a,3gp,3g2,mj2
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The language of the content. For more information, see the ISO 639-2 Alpha-3 codes for the representation of names of languages.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The coordinate pair of the central point. The coordinate pair consists of latitude and longitude values. This parameter value must be in the "latitude,longitude" format. Valid values of the latitude: [-90,+90]. Valid values of the longitude: [-180,+180].
	//
	// example:
	//
	// +120.029003,+30.283095
	LatLong *string `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	// The performer.
	//
	// example:
	//
	// unable
	Performer *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	// The time of recording. For more information about the time formats, see the RFC3339 Nano standard.
	//
	// example:
	//
	// 2022-04-24T02:39:57Z
	ProduceTime *string `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	// The number of programs.
	//
	// example:
	//
	// 2
	ProgramCount *int64 `json:"ProgramCount,omitempty" xml:"ProgramCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2213B1A9-EB3D-4666-84E0-24980BC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The size of the media object. Unit: bytes.
	//
	// example:
	//
	// 25115517
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The initial playback time.
	//
	// example:
	//
	// 0.000000
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of media streams.
	//
	// example:
	//
	// 2
	StreamCount *int64 `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	// The subtitle streams.
	Subtitles []*SubtitleStream `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The title of the media object.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The video height in pixels.
	//
	// example:
	//
	// 1920
	VideoHeight *int64 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// The video streams.
	VideoStreams []*VideoStream `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	// The video width in pixels.
	//
	// example:
	//
	// 1080
	VideoWidth *int64 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s DetectMediaMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectMediaMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaResponseBody) GetAddresses() []*Address {
	return s.Addresses
}

func (s *DetectMediaMetaResponseBody) GetAlbum() *string {
	return s.Album
}

func (s *DetectMediaMetaResponseBody) GetAlbumArtist() *string {
	return s.AlbumArtist
}

func (s *DetectMediaMetaResponseBody) GetArtist() *string {
	return s.Artist
}

func (s *DetectMediaMetaResponseBody) GetAudioStreams() []*AudioStream {
	return s.AudioStreams
}

func (s *DetectMediaMetaResponseBody) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *DetectMediaMetaResponseBody) GetComposer() *string {
	return s.Composer
}

func (s *DetectMediaMetaResponseBody) GetDuration() *float64 {
	return s.Duration
}

func (s *DetectMediaMetaResponseBody) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *DetectMediaMetaResponseBody) GetFormatName() *string {
	return s.FormatName
}

func (s *DetectMediaMetaResponseBody) GetLanguage() *string {
	return s.Language
}

func (s *DetectMediaMetaResponseBody) GetLatLong() *string {
	return s.LatLong
}

func (s *DetectMediaMetaResponseBody) GetPerformer() *string {
	return s.Performer
}

func (s *DetectMediaMetaResponseBody) GetProduceTime() *string {
	return s.ProduceTime
}

func (s *DetectMediaMetaResponseBody) GetProgramCount() *int64 {
	return s.ProgramCount
}

func (s *DetectMediaMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectMediaMetaResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *DetectMediaMetaResponseBody) GetStartTime() *float64 {
	return s.StartTime
}

func (s *DetectMediaMetaResponseBody) GetStreamCount() *int64 {
	return s.StreamCount
}

func (s *DetectMediaMetaResponseBody) GetSubtitles() []*SubtitleStream {
	return s.Subtitles
}

func (s *DetectMediaMetaResponseBody) GetTitle() *string {
	return s.Title
}

func (s *DetectMediaMetaResponseBody) GetVideoHeight() *int64 {
	return s.VideoHeight
}

func (s *DetectMediaMetaResponseBody) GetVideoStreams() []*VideoStream {
	return s.VideoStreams
}

func (s *DetectMediaMetaResponseBody) GetVideoWidth() *int64 {
	return s.VideoWidth
}

func (s *DetectMediaMetaResponseBody) SetAddresses(v []*Address) *DetectMediaMetaResponseBody {
	s.Addresses = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetAlbum(v string) *DetectMediaMetaResponseBody {
	s.Album = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetAlbumArtist(v string) *DetectMediaMetaResponseBody {
	s.AlbumArtist = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetArtist(v string) *DetectMediaMetaResponseBody {
	s.Artist = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetAudioStreams(v []*AudioStream) *DetectMediaMetaResponseBody {
	s.AudioStreams = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetBitrate(v int64) *DetectMediaMetaResponseBody {
	s.Bitrate = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetComposer(v string) *DetectMediaMetaResponseBody {
	s.Composer = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetDuration(v float64) *DetectMediaMetaResponseBody {
	s.Duration = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetFormatLongName(v string) *DetectMediaMetaResponseBody {
	s.FormatLongName = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetFormatName(v string) *DetectMediaMetaResponseBody {
	s.FormatName = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetLanguage(v string) *DetectMediaMetaResponseBody {
	s.Language = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetLatLong(v string) *DetectMediaMetaResponseBody {
	s.LatLong = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetPerformer(v string) *DetectMediaMetaResponseBody {
	s.Performer = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetProduceTime(v string) *DetectMediaMetaResponseBody {
	s.ProduceTime = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetProgramCount(v int64) *DetectMediaMetaResponseBody {
	s.ProgramCount = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetRequestId(v string) *DetectMediaMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetSize(v int64) *DetectMediaMetaResponseBody {
	s.Size = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetStartTime(v float64) *DetectMediaMetaResponseBody {
	s.StartTime = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetStreamCount(v int64) *DetectMediaMetaResponseBody {
	s.StreamCount = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetSubtitles(v []*SubtitleStream) *DetectMediaMetaResponseBody {
	s.Subtitles = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetTitle(v string) *DetectMediaMetaResponseBody {
	s.Title = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetVideoHeight(v int64) *DetectMediaMetaResponseBody {
	s.VideoHeight = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetVideoStreams(v []*VideoStream) *DetectMediaMetaResponseBody {
	s.VideoStreams = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetVideoWidth(v int64) *DetectMediaMetaResponseBody {
	s.VideoWidth = &v
	return s
}

func (s *DetectMediaMetaResponseBody) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AudioStreams != nil {
		for _, item := range s.AudioStreams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoStreams != nil {
		for _, item := range s.VideoStreams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
