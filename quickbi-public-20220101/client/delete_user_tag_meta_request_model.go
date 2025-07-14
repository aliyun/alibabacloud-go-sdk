// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserTagMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagId(v string) *DeleteUserTagMetaRequest
	GetTagId() *string
}

type DeleteUserTagMetaRequest struct {
	// The ID of the tag to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// pop_001
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeleteUserTagMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserTagMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserTagMetaRequest) GetTagId() *string {
	return s.TagId
}

func (s *DeleteUserTagMetaRequest) SetTagId(v string) *DeleteUserTagMetaRequest {
	s.TagId = &v
	return s
}

func (s *DeleteUserTagMetaRequest) Validate() error {
	return dara.Validate(s)
}
