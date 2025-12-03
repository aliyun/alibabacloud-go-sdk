// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemAllCommentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DeleteWorkitemAllCommentRequest
	GetIdentifier() *string
}

type DeleteWorkitemAllCommentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteWorkitemAllCommentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemAllCommentRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemAllCommentRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DeleteWorkitemAllCommentRequest) SetIdentifier(v string) *DeleteWorkitemAllCommentRequest {
	s.Identifier = &v
	return s
}

func (s *DeleteWorkitemAllCommentRequest) Validate() error {
	return dara.Validate(s)
}
