// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsForPrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListTagsForPrivateAccessApplicationRequest
	GetApplicationIds() []*string
}

type ListTagsForPrivateAccessApplicationRequest struct {
	// Collection of private network access application IDs. Enter a maximum of 100 private network access application IDs.
	//
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListTagsForPrivateAccessApplicationRequest) SetApplicationIds(v []*string) *ListTagsForPrivateAccessApplicationRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
