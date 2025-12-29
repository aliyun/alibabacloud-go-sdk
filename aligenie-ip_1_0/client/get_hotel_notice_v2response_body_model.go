// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetHotelNoticeV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelNoticeV2ResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelNoticeV2ResponseBodyResult) *GetHotelNoticeV2ResponseBody
	GetResult() *GetHotelNoticeV2ResponseBodyResult
	SetStatusCode(v int32) *GetHotelNoticeV2ResponseBody
	GetStatusCode() *int32
}

type GetHotelNoticeV2ResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0D0C***67DB
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelNoticeV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelNoticeV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelNoticeV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelNoticeV2ResponseBody) GetResult() *GetHotelNoticeV2ResponseBodyResult {
	return s.Result
}

func (s *GetHotelNoticeV2ResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelNoticeV2ResponseBody) SetMessage(v string) *GetHotelNoticeV2ResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) SetRequestId(v string) *GetHotelNoticeV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) SetResult(v *GetHotelNoticeV2ResponseBodyResult) *GetHotelNoticeV2ResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) SetStatusCode(v int32) *GetHotelNoticeV2ResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelNoticeV2ResponseBodyResult struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotelNoticeV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2ResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetHotelNoticeV2ResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelNoticeV2ResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetHotelNoticeV2ResponseBodyResult) SetContent(v string) *GetHotelNoticeV2ResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBodyResult) SetHotelId(v string) *GetHotelNoticeV2ResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBodyResult) SetTitle(v string) *GetHotelNoticeV2ResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetHotelNoticeV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
