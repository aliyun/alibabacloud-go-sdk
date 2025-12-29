// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelScreenSaverResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelScreenSaverResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelScreenSaverResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelScreenSaverResponseBodyResult) *GetHotelScreenSaverResponseBody
	GetResult() *GetHotelScreenSaverResponseBodyResult
}

type GetHotelScreenSaverResponseBody struct {
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
	// 5F0467E1-19F2-1757-B6D0-B79917BA2E81
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelScreenSaverResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelScreenSaverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelScreenSaverResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelScreenSaverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelScreenSaverResponseBody) GetResult() *GetHotelScreenSaverResponseBodyResult {
	return s.Result
}

func (s *GetHotelScreenSaverResponseBody) SetCode(v int32) *GetHotelScreenSaverResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetMessage(v string) *GetHotelScreenSaverResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetRequestId(v string) *GetHotelScreenSaverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelScreenSaverResponseBody) SetResult(v *GetHotelScreenSaverResponseBodyResult) *GetHotelScreenSaverResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelScreenSaverResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelScreenSaverResponseBodyResult struct {
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/3d4fe6d66ec49d9789635f66627f0339/welcome_audios/976210a6532150f49c2677a8b7dbc105/l6fspbhd.jpg
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	// example:
	//
	// common-weather
	StyleCode *string `json:"StyleCode,omitempty" xml:"StyleCode,omitempty"`
}

func (s GetHotelScreenSaverResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponseBodyResult) GetPicUrl() *string {
	return s.PicUrl
}

func (s *GetHotelScreenSaverResponseBodyResult) GetStyleCode() *string {
	return s.StyleCode
}

func (s *GetHotelScreenSaverResponseBodyResult) SetPicUrl(v string) *GetHotelScreenSaverResponseBodyResult {
	s.PicUrl = &v
	return s
}

func (s *GetHotelScreenSaverResponseBodyResult) SetStyleCode(v string) *GetHotelScreenSaverResponseBodyResult {
	s.StyleCode = &v
	return s
}

func (s *GetHotelScreenSaverResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
