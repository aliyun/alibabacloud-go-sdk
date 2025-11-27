// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceExistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsExistInstance(v bool) *CheckInstanceExistResponseBody
	GetIsExistInstance() *bool
	SetRequestId(v string) *CheckInstanceExistResponseBody
	GetRequestId() *string
}

type CheckInstanceExistResponseBody struct {
	IsExistInstance *bool   `json:"IsExistInstance,omitempty" xml:"IsExistInstance,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstanceExistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceExistResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceExistResponseBody) GetIsExistInstance() *bool {
	return s.IsExistInstance
}

func (s *CheckInstanceExistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceExistResponseBody) SetIsExistInstance(v bool) *CheckInstanceExistResponseBody {
	s.IsExistInstance = &v
	return s
}

func (s *CheckInstanceExistResponseBody) SetRequestId(v string) *CheckInstanceExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceExistResponseBody) Validate() error {
	return dara.Validate(s)
}
