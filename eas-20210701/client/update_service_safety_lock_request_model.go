// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSafetyLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLock(v string) *UpdateServiceSafetyLockRequest
	GetLock() *string
}

type UpdateServiceSafetyLockRequest struct {
	// The lock scope. Valid values:
	//
	// 	- all: locks all operations.
	//
	// 	- dangerous: locks dangerous operations such as delete and stop operations.
	//
	// 	- none: locks no operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// dangerous
	Lock *string `json:"Lock,omitempty" xml:"Lock,omitempty"`
}

func (s UpdateServiceSafetyLockRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSafetyLockRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceSafetyLockRequest) GetLock() *string {
	return s.Lock
}

func (s *UpdateServiceSafetyLockRequest) SetLock(v string) *UpdateServiceSafetyLockRequest {
	s.Lock = &v
	return s
}

func (s *UpdateServiceSafetyLockRequest) Validate() error {
	return dara.Validate(s)
}
