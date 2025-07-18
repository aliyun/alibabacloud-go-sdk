// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForPrivateAccessTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagIds(v []*string) *ListApplicationsForPrivateAccessTagRequest
	GetTagIds() []*string
}

type ListApplicationsForPrivateAccessTagRequest struct {
	// This parameter is required.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListApplicationsForPrivateAccessTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListApplicationsForPrivateAccessTagRequest) SetTagIds(v []*string) *ListApplicationsForPrivateAccessTagRequest {
	s.TagIds = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagRequest) Validate() error {
	return dara.Validate(s)
}
