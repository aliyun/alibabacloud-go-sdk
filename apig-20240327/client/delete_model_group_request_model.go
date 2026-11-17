// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpectedVersion(v int64) *DeleteModelGroupRequest
	GetExpectedVersion() *int64
}

type DeleteModelGroupRequest struct {
	// The expected version. Only this version is deleted. API saved-state references prevent deletion, and authorization references are cleared.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
}

func (s DeleteModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelGroupRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *DeleteModelGroupRequest) SetExpectedVersion(v int64) *DeleteModelGroupRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *DeleteModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
