// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v int32) *UpdateServiceVersionRequest
	GetVersion() *int32
}

type UpdateServiceVersionRequest struct {
	// The target version to which the service is switched. The value must be an integer greater than 0 and less than the current version of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRequest) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateServiceVersionRequest) SetVersion(v int32) *UpdateServiceVersionRequest {
	s.Version = &v
	return s
}

func (s *UpdateServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
