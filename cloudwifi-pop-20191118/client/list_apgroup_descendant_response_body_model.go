// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApgroupDescendantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *ListApgroupDescendantResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *ListApgroupDescendantResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *ListApgroupDescendantResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *ListApgroupDescendantResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListApgroupDescendantResponseBody
	GetRequestId() *string
}

type ListApgroupDescendantResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApgroupDescendantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApgroupDescendantResponseBody) GoString() string {
	return s.String()
}

func (s *ListApgroupDescendantResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ListApgroupDescendantResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *ListApgroupDescendantResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListApgroupDescendantResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListApgroupDescendantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApgroupDescendantResponseBody) SetData(v map[string]interface{}) *ListApgroupDescendantResponseBody {
	s.Data = v
	return s
}

func (s *ListApgroupDescendantResponseBody) SetErrorCode(v int32) *ListApgroupDescendantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListApgroupDescendantResponseBody) SetErrorMessage(v string) *ListApgroupDescendantResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListApgroupDescendantResponseBody) SetIsSuccess(v bool) *ListApgroupDescendantResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListApgroupDescendantResponseBody) SetRequestId(v string) *ListApgroupDescendantResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApgroupDescendantResponseBody) Validate() error {
	return dara.Validate(s)
}
