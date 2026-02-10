// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalLabelStudioServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteMultimodalLabelStudioServiceResponseBody
	GetData() interface{}
	SetErrCode(v string) *DeleteMultimodalLabelStudioServiceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteMultimodalLabelStudioServiceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteMultimodalLabelStudioServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMultimodalLabelStudioServiceResponseBody
	GetSuccess() *bool
}

type DeleteMultimodalLabelStudioServiceResponseBody struct {
	// example:
	//
	// 服务相关信息
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

func (s DeleteMultimodalLabelStudioServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalLabelStudioServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) SetData(v interface{}) *DeleteMultimodalLabelStudioServiceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) SetErrCode(v string) *DeleteMultimodalLabelStudioServiceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) SetErrMessage(v string) *DeleteMultimodalLabelStudioServiceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) SetRequestId(v string) *DeleteMultimodalLabelStudioServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) SetSuccess(v bool) *DeleteMultimodalLabelStudioServiceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
