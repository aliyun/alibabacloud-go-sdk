// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForPrivateAccessTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagIds(v []*string) *ListPolicesForPrivateAccessTagRequest
	GetTagIds() []*string
}

type ListPolicesForPrivateAccessTagRequest struct {
	// This parameter is required.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListPolicesForPrivateAccessTagRequest) SetTagIds(v []*string) *ListPolicesForPrivateAccessTagRequest {
	s.TagIds = v
	return s
}

func (s *ListPolicesForPrivateAccessTagRequest) Validate() error {
	return dara.Validate(s)
}
