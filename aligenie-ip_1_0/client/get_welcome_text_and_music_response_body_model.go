// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWelcomeTextAndMusicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *GetWelcomeTextAndMusicResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *GetWelcomeTextAndMusicResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWelcomeTextAndMusicResponseBody
	GetRequestId() *string
	SetResult(v *GetWelcomeTextAndMusicResponseBodyResult) *GetWelcomeTextAndMusicResponseBody
	GetResult() *GetWelcomeTextAndMusicResponseBodyResult
	SetStatusCode(v int32) *GetWelcomeTextAndMusicResponseBody
	GetStatusCode() *int32
}

type GetWelcomeTextAndMusicResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetWelcomeTextAndMusicResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetWelcomeTextAndMusicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWelcomeTextAndMusicResponseBody) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *GetWelcomeTextAndMusicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWelcomeTextAndMusicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWelcomeTextAndMusicResponseBody) GetResult() *GetWelcomeTextAndMusicResponseBodyResult {
	return s.Result
}

func (s *GetWelcomeTextAndMusicResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWelcomeTextAndMusicResponseBody) SetExtentions(v map[string]interface{}) *GetWelcomeTextAndMusicResponseBody {
	s.Extentions = v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetMessage(v string) *GetWelcomeTextAndMusicResponseBody {
	s.Message = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetRequestId(v string) *GetWelcomeTextAndMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetResult(v *GetWelcomeTextAndMusicResponseBodyResult) *GetWelcomeTextAndMusicResponseBody {
	s.Result = v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) SetStatusCode(v int32) *GetWelcomeTextAndMusicResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWelcomeTextAndMusicResponseBodyResult struct {
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// http://ailabsaicloudservice.alicdn.com/tmp/a.wav
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	Text     *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetWelcomeTextAndMusicResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetWelcomeTextAndMusicResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) GetText() *string {
	return s.Text
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) SetHotelId(v string) *GetWelcomeTextAndMusicResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) SetMusicUrl(v string) *GetWelcomeTextAndMusicResponseBodyResult {
	s.MusicUrl = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) SetText(v string) *GetWelcomeTextAndMusicResponseBodyResult {
	s.Text = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
