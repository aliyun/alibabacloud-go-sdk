// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelSourceModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpectedVersion(v int64) *DeleteAiModelSourceModelRequest
	GetExpectedVersion() *int64
}

type DeleteAiModelSourceModelRequest struct {
	// The expected object version. Use the version returned by the most recent query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
}

func (s DeleteAiModelSourceModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelSourceModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiModelSourceModelRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *DeleteAiModelSourceModelRequest) SetExpectedVersion(v int64) *DeleteAiModelSourceModelRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *DeleteAiModelSourceModelRequest) Validate() error {
	return dara.Validate(s)
}
