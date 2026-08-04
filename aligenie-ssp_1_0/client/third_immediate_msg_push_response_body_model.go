// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThirdImmediateMsgPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ThirdImmediateMsgPushResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ThirdImmediateMsgPushResponseBody
	GetErrorMsg() *string
	SetModel(v *ThirdImmediateMsgPushResponseBodyModel) *ThirdImmediateMsgPushResponseBody
	GetModel() *ThirdImmediateMsgPushResponseBodyModel
	SetSuccess(v bool) *ThirdImmediateMsgPushResponseBody
	GetSuccess() *bool
}

type ThirdImmediateMsgPushResponseBody struct {
	// error code
	//
	// example:
	//
	// 500000000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// error message
	//
	// example:
	//
	// 服务器内部异常
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// push result
	Model *ThirdImmediateMsgPushResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// whether the invocation succeeded
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ThirdImmediateMsgPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ThirdImmediateMsgPushResponseBody) GoString() string {
	return s.String()
}

func (s *ThirdImmediateMsgPushResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ThirdImmediateMsgPushResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ThirdImmediateMsgPushResponseBody) GetModel() *ThirdImmediateMsgPushResponseBodyModel {
	return s.Model
}

func (s *ThirdImmediateMsgPushResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ThirdImmediateMsgPushResponseBody) SetErrorCode(v string) *ThirdImmediateMsgPushResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ThirdImmediateMsgPushResponseBody) SetErrorMsg(v string) *ThirdImmediateMsgPushResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ThirdImmediateMsgPushResponseBody) SetModel(v *ThirdImmediateMsgPushResponseBodyModel) *ThirdImmediateMsgPushResponseBody {
	s.Model = v
	return s
}

func (s *ThirdImmediateMsgPushResponseBody) SetSuccess(v bool) *ThirdImmediateMsgPushResponseBody {
	s.Success = &v
	return s
}

func (s *ThirdImmediateMsgPushResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ThirdImmediateMsgPushResponseBodyModel struct {
	// request ID
	//
	// example:
	//
	// 2DF6FEFE-3301-16DD-ABCC-968A9524920B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// whether the push succeeded
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ThirdImmediateMsgPushResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ThirdImmediateMsgPushResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ThirdImmediateMsgPushResponseBodyModel) GetRequestId() *string {
	return s.RequestId
}

func (s *ThirdImmediateMsgPushResponseBodyModel) GetSuccess() *bool {
	return s.Success
}

func (s *ThirdImmediateMsgPushResponseBodyModel) SetRequestId(v string) *ThirdImmediateMsgPushResponseBodyModel {
	s.RequestId = &v
	return s
}

func (s *ThirdImmediateMsgPushResponseBodyModel) SetSuccess(v bool) *ThirdImmediateMsgPushResponseBodyModel {
	s.Success = &v
	return s
}

func (s *ThirdImmediateMsgPushResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
