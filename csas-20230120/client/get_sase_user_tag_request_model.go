// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaseUserTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagId(v string) *GetSaseUserTagRequest
	GetTagId() *string
}

type GetSaseUserTagRequest struct {
	// The user tag ID. You can obtain this value from the following operations:
	//
	// - [ListSaseUserTags](~~ListSaseUserTags~~): Lists user tags.
	//
	// - [CreateSaseUserTag](~~CreateSaseUserTag~~): Creates a user tag.
	//
	// example:
	//
	// su-tag-1ae52f66039fa0d4****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s GetSaseUserTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSaseUserTagRequest) GoString() string {
	return s.String()
}

func (s *GetSaseUserTagRequest) GetTagId() *string {
	return s.TagId
}

func (s *GetSaseUserTagRequest) SetTagId(v string) *GetSaseUserTagRequest {
	s.TagId = &v
	return s
}

func (s *GetSaseUserTagRequest) Validate() error {
	return dara.Validate(s)
}
