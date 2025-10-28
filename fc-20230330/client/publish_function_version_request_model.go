// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFunctionVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PublishVersionInput) *PublishFunctionVersionRequest
	GetBody() *PublishVersionInput
}

type PublishFunctionVersionRequest struct {
	// The information about the function version.
	//
	// This parameter is required.
	Body *PublishVersionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFunctionVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishFunctionVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishFunctionVersionRequest) GetBody() *PublishVersionInput {
	return s.Body
}

func (s *PublishFunctionVersionRequest) SetBody(v *PublishVersionInput) *PublishFunctionVersionRequest {
	s.Body = v
	return s
}

func (s *PublishFunctionVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
