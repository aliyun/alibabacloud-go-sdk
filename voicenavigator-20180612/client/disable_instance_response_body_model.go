// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DisableInstanceResponseBody
	GetStatus() *string
}

type DisableInstanceResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DisableInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DisableInstanceResponseBody) SetRequestId(v string) *DisableInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInstanceResponseBody) SetStatus(v string) *DisableInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DisableInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
