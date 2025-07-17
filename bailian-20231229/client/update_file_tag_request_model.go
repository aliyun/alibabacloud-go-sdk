// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTags(v []*string) *UpdateFileTagRequest
	GetTags() []*string
}

type UpdateFileTagRequest struct {
	// This parameter is required.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateFileTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileTagRequest) GetTags() []*string {
	return s.Tags
}

func (s *UpdateFileTagRequest) SetTags(v []*string) *UpdateFileTagRequest {
	s.Tags = v
	return s
}

func (s *UpdateFileTagRequest) Validate() error {
	return dara.Validate(s)
}
