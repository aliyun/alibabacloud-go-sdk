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
	// Indicates whether the instance exists. Valid values:
	//
	// - **true**: The instance exists.
	//
	// - **false**: The instance does not exist.
	//
	// example:
	//
	// true
	IsExistInstance *bool `json:"IsExistInstance,omitempty" xml:"IsExistInstance,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 11439B36-F703-49EB-8656-D3C87BE28B57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
