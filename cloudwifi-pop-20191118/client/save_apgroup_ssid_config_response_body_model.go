// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupSsidConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *SaveApgroupSsidConfigResponseBody
	GetData() []*int64
	SetErrorCode(v int32) *SaveApgroupSsidConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApgroupSsidConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApgroupSsidConfigResponseBody
	GetIsSuccess() *bool
}

type SaveApgroupSsidConfigResponseBody struct {
	Data         []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s SaveApgroupSsidConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupSsidConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApgroupSsidConfigResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *SaveApgroupSsidConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApgroupSsidConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApgroupSsidConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApgroupSsidConfigResponseBody) SetData(v []*int64) *SaveApgroupSsidConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApgroupSsidConfigResponseBody) SetErrorCode(v int32) *SaveApgroupSsidConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApgroupSsidConfigResponseBody) SetErrorMessage(v string) *SaveApgroupSsidConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApgroupSsidConfigResponseBody) SetIsSuccess(v bool) *SaveApgroupSsidConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApgroupSsidConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
