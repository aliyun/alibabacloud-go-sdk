// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDeviceStatusDetailResponseBody
	GetCode() *int32
	SetMessage(v string) *GetDeviceStatusDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceStatusDetailResponseBody
	GetRequestId() *string
	SetResult(v *GetDeviceStatusDetailResponseBodyResult) *GetDeviceStatusDetailResponseBody
	GetResult() *GetDeviceStatusDetailResponseBodyResult
}

type GetDeviceStatusDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceStatusDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceStatusDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDeviceStatusDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceStatusDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceStatusDetailResponseBody) GetResult() *GetDeviceStatusDetailResponseBodyResult {
	return s.Result
}

func (s *GetDeviceStatusDetailResponseBody) SetCode(v int32) *GetDeviceStatusDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) SetMessage(v string) *GetDeviceStatusDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) SetRequestId(v string) *GetDeviceStatusDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) SetResult(v *GetDeviceStatusDetailResponseBodyResult) *GetDeviceStatusDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceStatusDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceStatusDetailResponseBodyResult struct {
	Player  *GetDeviceStatusDetailResponseBodyResultPlayer  `json:"Player,omitempty" xml:"Player,omitempty" type:"Struct"`
	Power   *GetDeviceStatusDetailResponseBodyResultPower   `json:"Power,omitempty" xml:"Power,omitempty" type:"Struct"`
	Speaker *GetDeviceStatusDetailResponseBodyResultSpeaker `json:"Speaker,omitempty" xml:"Speaker,omitempty" type:"Struct"`
}

func (s GetDeviceStatusDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResult) GetPlayer() *GetDeviceStatusDetailResponseBodyResultPlayer {
	return s.Player
}

func (s *GetDeviceStatusDetailResponseBodyResult) GetPower() *GetDeviceStatusDetailResponseBodyResultPower {
	return s.Power
}

func (s *GetDeviceStatusDetailResponseBodyResult) GetSpeaker() *GetDeviceStatusDetailResponseBodyResultSpeaker {
	return s.Speaker
}

func (s *GetDeviceStatusDetailResponseBodyResult) SetPlayer(v *GetDeviceStatusDetailResponseBodyResultPlayer) *GetDeviceStatusDetailResponseBodyResult {
	s.Player = v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResult) SetPower(v *GetDeviceStatusDetailResponseBodyResultPower) *GetDeviceStatusDetailResponseBodyResult {
	s.Power = v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResult) SetSpeaker(v *GetDeviceStatusDetailResponseBodyResultSpeaker) *GetDeviceStatusDetailResponseBodyResult {
	s.Speaker = v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetDeviceStatusDetailResponseBodyResultPlayer struct {
	AudioAlbum *string `json:"AudioAlbum,omitempty" xml:"AudioAlbum,omitempty"`
	// example:
	//
	// 刘德华
	AudioAnchor *string `json:"AudioAnchor,omitempty" xml:"AudioAnchor,omitempty"`
	// example:
	//
	// ""
	AudioExt *string `json:"AudioExt,omitempty" xml:"AudioExt,omitempty"`
	// example:
	//
	// 123
	AudioId *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	// example:
	//
	// 253
	AudioLength *string `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	AudioName   *string `json:"AudioName,omitempty" xml:"AudioName,omitempty"`
	// example:
	//
	// xiami
	AudioSource *string `json:"AudioSource,omitempty" xml:"AudioSource,omitempty"`
	// example:
	//
	// https://xxx
	AudioUrl *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	// example:
	//
	// mp3
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 30
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// cloud
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// pause
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetDeviceStatusDetailResponseBodyResultPlayer) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResultPlayer) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioAlbum() *string {
	return s.AudioAlbum
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioAnchor() *string {
	return s.AudioAnchor
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioExt() *string {
	return s.AudioExt
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioId() *string {
	return s.AudioId
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioLength() *string {
	return s.AudioLength
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioName() *string {
	return s.AudioName
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioSource() *string {
	return s.AudioSource
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetAudioUrl() *string {
	return s.AudioUrl
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetFormat() *string {
	return s.Format
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetProgress() *string {
	return s.Progress
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetSource() *string {
	return s.Source
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetStatus() *string {
	return s.Status
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioAlbum(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioAlbum = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioAnchor(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioAnchor = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioExt(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioExt = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioId(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioId = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioLength(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioLength = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioName(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioName = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioSource(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioSource = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetAudioUrl(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.AudioUrl = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetFormat(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Format = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetProgress(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Progress = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetSource(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Source = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetStatus(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Status = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) SetTimestamp(v string) *GetDeviceStatusDetailResponseBodyResultPlayer {
	s.Timestamp = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPlayer) Validate() error {
	return dara.Validate(s)
}

type GetDeviceStatusDetailResponseBodyResultPower struct {
	// example:
	//
	// 30
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceStatusDetailResponseBodyResultPower) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResultPower) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) GetStatus() *string {
	return s.Status
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) SetQuantity(v int32) *GetDeviceStatusDetailResponseBodyResultPower {
	s.Quantity = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) SetStatus(v string) *GetDeviceStatusDetailResponseBodyResultPower {
	s.Status = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultPower) Validate() error {
	return dara.Validate(s)
}

type GetDeviceStatusDetailResponseBodyResultSpeaker struct {
	Muted *bool `json:"Muted,omitempty" xml:"Muted,omitempty"`
	// example:
	//
	// 10
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetDeviceStatusDetailResponseBodyResultSpeaker) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailResponseBodyResultSpeaker) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) GetMuted() *bool {
	return s.Muted
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) GetVolume() *int32 {
	return s.Volume
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) SetMuted(v bool) *GetDeviceStatusDetailResponseBodyResultSpeaker {
	s.Muted = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) SetVolume(v int32) *GetDeviceStatusDetailResponseBodyResultSpeaker {
	s.Volume = &v
	return s
}

func (s *GetDeviceStatusDetailResponseBodyResultSpeaker) Validate() error {
	return dara.Validate(s)
}
