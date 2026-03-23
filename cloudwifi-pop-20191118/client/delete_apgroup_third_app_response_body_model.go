// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupThirdAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *DeleteApgroupThirdAppResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *DeleteApgroupThirdAppResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *DeleteApgroupThirdAppResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *DeleteApgroupThirdAppResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteApgroupThirdAppResponseBody
	GetRequestId() *string
}

type DeleteApgroupThirdAppResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApgroupThirdAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupThirdAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApgroupThirdAppResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *DeleteApgroupThirdAppResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DeleteApgroupThirdAppResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteApgroupThirdAppResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteApgroupThirdAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApgroupThirdAppResponseBody) SetData(v map[string]interface{}) *DeleteApgroupThirdAppResponseBody {
	s.Data = v
	return s
}

func (s *DeleteApgroupThirdAppResponseBody) SetErrorCode(v int32) *DeleteApgroupThirdAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApgroupThirdAppResponseBody) SetErrorMessage(v string) *DeleteApgroupThirdAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteApgroupThirdAppResponseBody) SetIsSuccess(v bool) *DeleteApgroupThirdAppResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteApgroupThirdAppResponseBody) SetRequestId(v string) *DeleteApgroupThirdAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApgroupThirdAppResponseBody) Validate() error {
	return dara.Validate(s)
}
