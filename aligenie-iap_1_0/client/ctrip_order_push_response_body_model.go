// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCtripOrderPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CtripOrderPushResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CtripOrderPushResponseBody
	GetErrorMsg() *string
	SetModel(v *CtripOrderPushResponseBodyModel) *CtripOrderPushResponseBody
	GetModel() *CtripOrderPushResponseBodyModel
	SetSuccess(v bool) *CtripOrderPushResponseBody
	GetSuccess() *bool
}

type CtripOrderPushResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// mdoel
	Model *CtripOrderPushResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CtripOrderPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CtripOrderPushResponseBody) GoString() string {
	return s.String()
}

func (s *CtripOrderPushResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CtripOrderPushResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CtripOrderPushResponseBody) GetModel() *CtripOrderPushResponseBodyModel {
	return s.Model
}

func (s *CtripOrderPushResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CtripOrderPushResponseBody) SetErrorCode(v string) *CtripOrderPushResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CtripOrderPushResponseBody) SetErrorMsg(v string) *CtripOrderPushResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CtripOrderPushResponseBody) SetModel(v *CtripOrderPushResponseBodyModel) *CtripOrderPushResponseBody {
	s.Model = v
	return s
}

func (s *CtripOrderPushResponseBody) SetSuccess(v bool) *CtripOrderPushResponseBody {
	s.Success = &v
	return s
}

func (s *CtripOrderPushResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CtripOrderPushResponseBodyModel struct {
	// example:
	//
	// 1234566789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CtripOrderPushResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CtripOrderPushResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CtripOrderPushResponseBodyModel) GetRequestId() *string {
	return s.RequestId
}

func (s *CtripOrderPushResponseBodyModel) GetSuccess() *bool {
	return s.Success
}

func (s *CtripOrderPushResponseBodyModel) SetRequestId(v string) *CtripOrderPushResponseBodyModel {
	s.RequestId = &v
	return s
}

func (s *CtripOrderPushResponseBodyModel) SetSuccess(v bool) *CtripOrderPushResponseBodyModel {
	s.Success = &v
	return s
}

func (s *CtripOrderPushResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
