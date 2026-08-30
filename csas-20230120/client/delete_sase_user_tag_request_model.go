// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaseUserTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagIds(v []*string) *DeleteSaseUserTagRequest
	GetTagIds() []*string
}

type DeleteSaseUserTagRequest struct {
	// The collection of user tag IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s DeleteSaseUserTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaseUserTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteSaseUserTagRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *DeleteSaseUserTagRequest) SetTagIds(v []*string) *DeleteSaseUserTagRequest {
	s.TagIds = v
	return s
}

func (s *DeleteSaseUserTagRequest) Validate() error {
	return dara.Validate(s)
}
