// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagOptionId(v string) *DeleteTagOptionRequest
	GetTagOptionId() *string
}

type DeleteTagOptionRequest struct {
	// The ID of the tag option.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-bp1u6mdf3d****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s DeleteTagOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagOptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagOptionRequest) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *DeleteTagOptionRequest) SetTagOptionId(v string) *DeleteTagOptionRequest {
	s.TagOptionId = &v
	return s
}

func (s *DeleteTagOptionRequest) Validate() error {
	return dara.Validate(s)
}
