// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalLabelStudioServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *CreateMultimodalLabelStudioServiceResponseBody
	GetData() interface{}
	SetErrCode(v string) *CreateMultimodalLabelStudioServiceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMultimodalLabelStudioServiceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateMultimodalLabelStudioServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMultimodalLabelStudioServiceResponseBody
	GetSuccess() *bool
}

type CreateMultimodalLabelStudioServiceResponseBody struct {
	// example:
	//
	// JsonObject
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

func (s CreateMultimodalLabelStudioServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalLabelStudioServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) SetData(v interface{}) *CreateMultimodalLabelStudioServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) SetErrCode(v string) *CreateMultimodalLabelStudioServiceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) SetErrMessage(v string) *CreateMultimodalLabelStudioServiceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) SetRequestId(v string) *CreateMultimodalLabelStudioServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) SetSuccess(v bool) *CreateMultimodalLabelStudioServiceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
