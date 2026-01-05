// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagOptionId(v string) *GetTagOptionRequest
	GetTagOptionId() *string
}

type GetTagOptionRequest struct {
	// The ID of the tag option.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-bp1r3mxq3t****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s GetTagOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTagOptionRequest) GoString() string {
	return s.String()
}

func (s *GetTagOptionRequest) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *GetTagOptionRequest) SetTagOptionId(v string) *GetTagOptionRequest {
	s.TagOptionId = &v
	return s
}

func (s *GetTagOptionRequest) Validate() error {
	return dara.Validate(s)
}
