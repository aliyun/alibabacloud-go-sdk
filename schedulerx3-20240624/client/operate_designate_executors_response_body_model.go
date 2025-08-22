// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDesignateExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateDesignateExecutorsResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateDesignateExecutorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateDesignateExecutorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateDesignateExecutorsResponseBody
	GetSuccess() *bool
}

type OperateDesignateExecutorsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AFD5B166-4A7D-50DF-91BF-EFAFD41F7335
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateDesignateExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateDesignateExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateDesignateExecutorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateDesignateExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateDesignateExecutorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateDesignateExecutorsResponseBody) SetCode(v int32) *OperateDesignateExecutorsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) SetMessage(v string) *OperateDesignateExecutorsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) SetRequestId(v string) *OperateDesignateExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) SetSuccess(v bool) *OperateDesignateExecutorsResponseBody {
	s.Success = &v
	return s
}

func (s *OperateDesignateExecutorsResponseBody) Validate() error {
	return dara.Validate(s)
}
