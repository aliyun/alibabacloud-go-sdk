// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoPlaylistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioPlaylist(v []*GenerateVideoPlaylistResponseBodyAudioPlaylist) *GenerateVideoPlaylistResponseBody
	GetAudioPlaylist() []*GenerateVideoPlaylistResponseBodyAudioPlaylist
	SetDuration(v float32) *GenerateVideoPlaylistResponseBody
	GetDuration() *float32
	SetMasterURI(v string) *GenerateVideoPlaylistResponseBody
	GetMasterURI() *string
	SetRequestId(v string) *GenerateVideoPlaylistResponseBody
	GetRequestId() *string
	SetSubtitlePlaylist(v []*GenerateVideoPlaylistResponseBodySubtitlePlaylist) *GenerateVideoPlaylistResponseBody
	GetSubtitlePlaylist() []*GenerateVideoPlaylistResponseBodySubtitlePlaylist
	SetToken(v string) *GenerateVideoPlaylistResponseBody
	GetToken() *string
	SetVideoPlaylist(v []*GenerateVideoPlaylistResponseBodyVideoPlaylist) *GenerateVideoPlaylistResponseBody
	GetVideoPlaylist() []*GenerateVideoPlaylistResponseBodyVideoPlaylist
}

type GenerateVideoPlaylistResponseBody struct {
	// The list of audio Media Playlist files.
	AudioPlaylist []*GenerateVideoPlaylistResponseBodyAudioPlaylist `json:"AudioPlaylist,omitempty" xml:"AudioPlaylist,omitempty" type:"Repeated"`
	// The total duration of the output video.
	//
	// example:
	//
	// 1082
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The OSS URI of the Master Playlist.
	//
	// example:
	//
	// oss://test-bucket/test-object/master.m3u8
	MasterURI *string `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFF*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of subtitle Media Playlist files.
	SubtitlePlaylist []*GenerateVideoPlaylistResponseBodySubtitlePlaylist `json:"SubtitlePlaylist,omitempty" xml:"SubtitlePlaylist,omitempty" type:"Repeated"`
	// The token of the Master Playlist.
	//
	// example:
	//
	// 92376fbb-171f-4259-913f-705f7ee0****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The list of video Media Playlist files.
	VideoPlaylist []*GenerateVideoPlaylistResponseBodyVideoPlaylist `json:"VideoPlaylist,omitempty" xml:"VideoPlaylist,omitempty" type:"Repeated"`
}

func (s GenerateVideoPlaylistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBody) GetAudioPlaylist() []*GenerateVideoPlaylistResponseBodyAudioPlaylist {
	return s.AudioPlaylist
}

func (s *GenerateVideoPlaylistResponseBody) GetDuration() *float32 {
	return s.Duration
}

func (s *GenerateVideoPlaylistResponseBody) GetMasterURI() *string {
	return s.MasterURI
}

func (s *GenerateVideoPlaylistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateVideoPlaylistResponseBody) GetSubtitlePlaylist() []*GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	return s.SubtitlePlaylist
}

func (s *GenerateVideoPlaylistResponseBody) GetToken() *string {
	return s.Token
}

func (s *GenerateVideoPlaylistResponseBody) GetVideoPlaylist() []*GenerateVideoPlaylistResponseBodyVideoPlaylist {
	return s.VideoPlaylist
}

func (s *GenerateVideoPlaylistResponseBody) SetAudioPlaylist(v []*GenerateVideoPlaylistResponseBodyAudioPlaylist) *GenerateVideoPlaylistResponseBody {
	s.AudioPlaylist = v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetDuration(v float32) *GenerateVideoPlaylistResponseBody {
	s.Duration = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetMasterURI(v string) *GenerateVideoPlaylistResponseBody {
	s.MasterURI = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetRequestId(v string) *GenerateVideoPlaylistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetSubtitlePlaylist(v []*GenerateVideoPlaylistResponseBodySubtitlePlaylist) *GenerateVideoPlaylistResponseBody {
	s.SubtitlePlaylist = v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetToken(v string) *GenerateVideoPlaylistResponseBody {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetVideoPlaylist(v []*GenerateVideoPlaylistResponseBodyVideoPlaylist) *GenerateVideoPlaylistResponseBody {
	s.VideoPlaylist = v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) Validate() error {
	if s.AudioPlaylist != nil {
		for _, item := range s.AudioPlaylist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubtitlePlaylist != nil {
		for _, item := range s.SubtitlePlaylist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoPlaylist != nil {
		for _, item := range s.VideoPlaylist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateVideoPlaylistResponseBodyAudioPlaylist struct {
	// The number of audio channels.
	//
	// example:
	//
	// 1
	Channels *int32 `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The token generated for the audio Media Playlist. You can use this parameter to construct the URI of the generated TS file.
	//
	// example:
	//
	// affe0c6042f09722fec95a21b8b******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The OSS URI of the audio Media Playlist.
	//
	// example:
	//
	// oss://test-bucket/test-object/output-audio.m3u8
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistResponseBodyAudioPlaylist) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBodyAudioPlaylist) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) GetChannels() *int32 {
	return s.Channels
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) GetToken() *string {
	return s.Token
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) GetURI() *string {
	return s.URI
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) SetChannels(v int32) *GenerateVideoPlaylistResponseBodyAudioPlaylist {
	s.Channels = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) SetToken(v string) *GenerateVideoPlaylistResponseBodyAudioPlaylist {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) SetURI(v string) *GenerateVideoPlaylistResponseBodyAudioPlaylist {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) Validate() error {
	return dara.Validate(s)
}

type GenerateVideoPlaylistResponseBodySubtitlePlaylist struct {
	// The sequence number of the subtitle stream, starting from 0.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language of the subtitle stream.
	//
	// > The language is obtained from the subtitle stream information of the source video specified by SourceURI. If the source video does not contain language information, this parameter is empty.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The token generated for the subtitle Media Playlist. You can use this parameter to construct the URI of the generated subtitle file.
	//
	// > You can use the returned token value to construct the URI of the transcoded subtitle file. The format is oss\\://${Bucket}/${Object}-${Token}_${Index}.ts. oss\\://${Bucket}/${Object} is the subtitle URI specified in the request parameters. ${Token} is the returned parameter. ${Index} is the sequence number of the subtitle.
	//
	// example:
	//
	// affe0c6042f09722fec95a21b8b******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The OSS URI of the subtitle Media Playlist.
	//
	// example:
	//
	// oss://test-bucket/test-object/output-subtitle.m3u8
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistResponseBodySubtitlePlaylist) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBodySubtitlePlaylist) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) GetIndex() *int32 {
	return s.Index
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) GetLanguage() *string {
	return s.Language
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) GetToken() *string {
	return s.Token
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) GetURI() *string {
	return s.URI
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetIndex(v int32) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.Index = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetLanguage(v string) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.Language = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetToken(v string) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetURI(v string) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) Validate() error {
	return dara.Validate(s)
}

type GenerateVideoPlaylistResponseBodyVideoPlaylist struct {
	// The video frame rate.
	//
	// example:
	//
	// 25/1
	FrameRate *string `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// The video resolution.
	//
	// example:
	//
	// 640x480
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The token generated for the video Media Playlist. You can use this parameter to construct the URI of the generated TS file.
	//
	// > You can use the returned token value to construct the URI of the transcoded TS file. The format is oss\\://${Bucket}/${Object}-${Token}-${Index}.ts. oss\\://${Bucket}/${Object} is the target URI specified in the request parameters. ${Token} is the returned parameter. ${Index} is the sequence number of the TS file.
	//
	// example:
	//
	// affe0c6042f09722fec95a21b8b******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The OSS URI of the video Media Playlist.
	//
	// example:
	//
	// oss://test-bucket/test-object/output-video.m3u8
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistResponseBodyVideoPlaylist) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBodyVideoPlaylist) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) GetFrameRate() *string {
	return s.FrameRate
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) GetResolution() *string {
	return s.Resolution
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) GetToken() *string {
	return s.Token
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) GetURI() *string {
	return s.URI
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) SetFrameRate(v string) *GenerateVideoPlaylistResponseBodyVideoPlaylist {
	s.FrameRate = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) SetResolution(v string) *GenerateVideoPlaylistResponseBodyVideoPlaylist {
	s.Resolution = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) SetToken(v string) *GenerateVideoPlaylistResponseBodyVideoPlaylist {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) SetURI(v string) *GenerateVideoPlaylistResponseBodyVideoPlaylist {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) Validate() error {
	return dara.Validate(s)
}
