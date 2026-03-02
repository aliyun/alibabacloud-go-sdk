// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ResourceCreateCmd) *CreateResourceRequest
	GetBody() *ResourceCreateCmd
}

type CreateResourceRequest struct {
	// This parameter is required.
	Body *ResourceCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) GetBody() *ResourceCreateCmd {
	return s.Body
}

func (s *CreateResourceRequest) SetBody(v *ResourceCreateCmd) *CreateResourceRequest {
	s.Body = v
	return s
}

func (s *CreateResourceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
