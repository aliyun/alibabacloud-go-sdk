// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupSsidConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []map[string]interface{}) *GetApgroupSsidConfigResponseBody
	GetData() []map[string]interface{}
	SetErrorCode(v int32) *GetApgroupSsidConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApgroupSsidConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApgroupSsidConfigResponseBody
	GetIsSuccess() *bool
}

type GetApgroupSsidConfigResponseBody struct {
	Data         []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApgroupSsidConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupSsidConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApgroupSsidConfigResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *GetApgroupSsidConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApgroupSsidConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApgroupSsidConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApgroupSsidConfigResponseBody) SetData(v []map[string]interface{}) *GetApgroupSsidConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetApgroupSsidConfigResponseBody) SetErrorCode(v int32) *GetApgroupSsidConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApgroupSsidConfigResponseBody) SetErrorMessage(v string) *GetApgroupSsidConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApgroupSsidConfigResponseBody) SetIsSuccess(v bool) *GetApgroupSsidConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApgroupSsidConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
