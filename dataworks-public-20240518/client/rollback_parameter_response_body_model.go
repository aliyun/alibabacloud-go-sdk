// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackParameterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RollbackParameterResponseBody
	GetSuccess() *bool
}

type RollbackParameterResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackParameterResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackParameterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackParameterResponseBody) SetRequestId(v string) *RollbackParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackParameterResponseBody) SetSuccess(v bool) *RollbackParameterResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
