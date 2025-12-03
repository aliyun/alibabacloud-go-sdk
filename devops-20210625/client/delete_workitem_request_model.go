// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DeleteWorkitemRequest
	GetIdentifier() *string
}

type DeleteWorkitemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3ef2ef6018d254d660e65f87a6
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteWorkitemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DeleteWorkitemRequest) SetIdentifier(v string) *DeleteWorkitemRequest {
	s.Identifier = &v
	return s
}

func (s *DeleteWorkitemRequest) Validate() error {
	return dara.Validate(s)
}
