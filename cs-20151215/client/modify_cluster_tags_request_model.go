// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*Tag) *ModifyClusterTagsRequest
	GetBody() []*Tag
}

type ModifyClusterTagsRequest struct {
	// The data of the tags that you want to modify.
	Body []*Tag `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ModifyClusterTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterTagsRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsRequest) GetBody() []*Tag {
	return s.Body
}

func (s *ModifyClusterTagsRequest) SetBody(v []*Tag) *ModifyClusterTagsRequest {
	s.Body = v
	return s
}

func (s *ModifyClusterTagsRequest) Validate() error {
	return dara.Validate(s)
}
