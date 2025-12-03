// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemCommentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommentId(v int64) *DeleteWorkitemCommentRequest
	GetCommentId() *int64
	SetIdentifier(v string) *DeleteWorkitemCommentRequest
	GetIdentifier() *string
}

type DeleteWorkitemCommentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12xx34
	CommentId *int64 `json:"commentId,omitempty" xml:"commentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e8b2xxxxxx2abdxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteWorkitemCommentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemCommentRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemCommentRequest) GetCommentId() *int64 {
	return s.CommentId
}

func (s *DeleteWorkitemCommentRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DeleteWorkitemCommentRequest) SetCommentId(v int64) *DeleteWorkitemCommentRequest {
	s.CommentId = &v
	return s
}

func (s *DeleteWorkitemCommentRequest) SetIdentifier(v string) *DeleteWorkitemCommentRequest {
	s.Identifier = &v
	return s
}

func (s *DeleteWorkitemCommentRequest) Validate() error {
	return dara.Validate(s)
}
