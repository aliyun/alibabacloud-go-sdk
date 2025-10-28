// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *TagResourcesInput) *TagResourcesRequest
	GetBody() *TagResourcesInput
}

type TagResourcesRequest struct {
	// The configuration of the resource tag.
	//
	// This parameter is required.
	Body *TagResourcesInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetBody() *TagResourcesInput {
	return s.Body
}

func (s *TagResourcesRequest) SetBody(v *TagResourcesInput) *TagResourcesRequest {
	s.Body = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
