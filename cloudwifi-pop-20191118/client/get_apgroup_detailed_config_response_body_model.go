// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupDetailedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApgroupDetailedConfigResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApgroupDetailedConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApgroupDetailedConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApgroupDetailedConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetApgroupDetailedConfigResponseBody
	GetRequestId() *string
}

type GetApgroupDetailedConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApgroupDetailedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupDetailedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApgroupDetailedConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApgroupDetailedConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApgroupDetailedConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApgroupDetailedConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApgroupDetailedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApgroupDetailedConfigResponseBody) SetData(v map[string]interface{}) *GetApgroupDetailedConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetApgroupDetailedConfigResponseBody) SetErrorCode(v int32) *GetApgroupDetailedConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApgroupDetailedConfigResponseBody) SetErrorMessage(v string) *GetApgroupDetailedConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApgroupDetailedConfigResponseBody) SetIsSuccess(v bool) *GetApgroupDetailedConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApgroupDetailedConfigResponseBody) SetRequestId(v string) *GetApgroupDetailedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApgroupDetailedConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
