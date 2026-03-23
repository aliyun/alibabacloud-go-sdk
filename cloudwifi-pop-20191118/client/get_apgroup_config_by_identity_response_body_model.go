// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupConfigByIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApgroupConfigByIdentityResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApgroupConfigByIdentityResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApgroupConfigByIdentityResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApgroupConfigByIdentityResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetApgroupConfigByIdentityResponseBody
	GetRequestId() *string
}

type GetApgroupConfigByIdentityResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApgroupConfigByIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupConfigByIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetApgroupConfigByIdentityResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApgroupConfigByIdentityResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApgroupConfigByIdentityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApgroupConfigByIdentityResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApgroupConfigByIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApgroupConfigByIdentityResponseBody) SetData(v map[string]interface{}) *GetApgroupConfigByIdentityResponseBody {
	s.Data = v
	return s
}

func (s *GetApgroupConfigByIdentityResponseBody) SetErrorCode(v int32) *GetApgroupConfigByIdentityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApgroupConfigByIdentityResponseBody) SetErrorMessage(v string) *GetApgroupConfigByIdentityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApgroupConfigByIdentityResponseBody) SetIsSuccess(v bool) *GetApgroupConfigByIdentityResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApgroupConfigByIdentityResponseBody) SetRequestId(v string) *GetApgroupConfigByIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApgroupConfigByIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
