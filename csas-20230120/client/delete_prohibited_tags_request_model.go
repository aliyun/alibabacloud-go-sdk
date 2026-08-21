// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagIds(v []*string) *DeleteProhibitedTagsRequest
	GetTagIds() []*string
}

type DeleteProhibitedTagsRequest struct {
	// This parameter is required.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s DeleteProhibitedTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedTagsRequest) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedTagsRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *DeleteProhibitedTagsRequest) SetTagIds(v []*string) *DeleteProhibitedTagsRequest {
	s.TagIds = v
	return s
}

func (s *DeleteProhibitedTagsRequest) Validate() error {
	return dara.Validate(s)
}
