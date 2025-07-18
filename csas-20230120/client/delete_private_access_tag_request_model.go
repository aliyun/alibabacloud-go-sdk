// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagId(v string) *DeletePrivateAccessTagRequest
	GetTagId() *string
}

type DeletePrivateAccessTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeletePrivateAccessTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessTagRequest) GetTagId() *string {
	return s.TagId
}

func (s *DeletePrivateAccessTagRequest) SetTagId(v string) *DeletePrivateAccessTagRequest {
	s.TagId = &v
	return s
}

func (s *DeletePrivateAccessTagRequest) Validate() error {
	return dara.Validate(s)
}
