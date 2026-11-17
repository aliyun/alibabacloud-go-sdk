// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpectedVersion(v int64) *DeleteAiModelSourceRequest
	GetExpectedVersion() *int64
}

type DeleteAiModelSourceRequest struct {
	// The expected object version. Use the version returned by the most recent query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
}

func (s DeleteAiModelSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiModelSourceRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *DeleteAiModelSourceRequest) SetExpectedVersion(v int64) *DeleteAiModelSourceRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *DeleteAiModelSourceRequest) Validate() error {
	return dara.Validate(s)
}
