// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicalModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpectedRevision(v int64) *DeleteLogicalModelRequest
	GetExpectedRevision() *int64
}

type DeleteLogicalModelRequest struct {
	// The expected object version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedRevision *int64 `json:"expectedRevision,omitempty" xml:"expectedRevision,omitempty"`
}

func (s DeleteLogicalModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicalModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogicalModelRequest) GetExpectedRevision() *int64 {
	return s.ExpectedRevision
}

func (s *DeleteLogicalModelRequest) SetExpectedRevision(v int64) *DeleteLogicalModelRequest {
	s.ExpectedRevision = &v
	return s
}

func (s *DeleteLogicalModelRequest) Validate() error {
	return dara.Validate(s)
}
