// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateApplicationResponseBody
	GetRequestId() *string
}

type OperateApplicationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *OperateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateApplicationResponseBody) SetRequestId(v string) *OperateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
