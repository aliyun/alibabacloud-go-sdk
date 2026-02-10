// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalLabelStudioServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalLabelStudioServiceResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalLabelStudioServiceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalLabelStudioServiceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalLabelStudioServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalLabelStudioServiceResponseBody
	GetSuccess() *bool
}

type ListMultimodalLabelStudioServiceResponseBody struct {
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMultimodalLabelStudioServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalLabelStudioServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalLabelStudioServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalLabelStudioServiceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalLabelStudioServiceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalLabelStudioServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalLabelStudioServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalLabelStudioServiceResponseBody) SetData(v interface{}) *ListMultimodalLabelStudioServiceResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponseBody) SetErrCode(v string) *ListMultimodalLabelStudioServiceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponseBody) SetErrMessage(v string) *ListMultimodalLabelStudioServiceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponseBody) SetRequestId(v string) *ListMultimodalLabelStudioServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponseBody) SetSuccess(v bool) *ListMultimodalLabelStudioServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
