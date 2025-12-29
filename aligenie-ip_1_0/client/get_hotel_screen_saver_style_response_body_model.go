// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetHotelScreenSaverStyleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelScreenSaverStyleResponseBody
	GetRequestId() *string
	SetResult(v []*GetHotelScreenSaverStyleResponseBodyResult) *GetHotelScreenSaverStyleResponseBody
	GetResult() []*GetHotelScreenSaverStyleResponseBodyResult
	SetStatusCode(v int32) *GetHotelScreenSaverStyleResponseBody
	GetStatusCode() *int32
}

type GetHotelScreenSaverStyleResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67**6FA
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelScreenSaverStyleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelScreenSaverStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelScreenSaverStyleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelScreenSaverStyleResponseBody) GetResult() []*GetHotelScreenSaverStyleResponseBodyResult {
	return s.Result
}

func (s *GetHotelScreenSaverStyleResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelScreenSaverStyleResponseBody) SetMessage(v string) *GetHotelScreenSaverStyleResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) SetRequestId(v string) *GetHotelScreenSaverStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) SetResult(v []*GetHotelScreenSaverStyleResponseBodyResult) *GetHotelScreenSaverStyleResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) SetStatusCode(v int32) *GetHotelScreenSaverStyleResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotelScreenSaverStyleResponseBodyResult struct {
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// example:
	//
	// common-weather
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// common-weather
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// example:
	//
	// https://img.***.png
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s GetHotelScreenSaverStyleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverStyleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) GetCnName() *string {
	return s.CnName
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) GetCode() *string {
	return s.Code
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) GetEnName() *string {
	return s.EnName
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) GetPicUrl() *string {
	return s.PicUrl
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetCnName(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.CnName = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetCode(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.Code = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetEnName(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.EnName = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) SetPicUrl(v string) *GetHotelScreenSaverStyleResponseBodyResult {
	s.PicUrl = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
