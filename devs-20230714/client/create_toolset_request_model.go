// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolsetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Toolset) *CreateToolsetRequest
	GetBody() *Toolset
}

type CreateToolsetRequest struct {
	Body *Toolset `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateToolsetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateToolsetRequest) GoString() string {
	return s.String()
}

func (s *CreateToolsetRequest) GetBody() *Toolset {
	return s.Body
}

func (s *CreateToolsetRequest) SetBody(v *Toolset) *CreateToolsetRequest {
	s.Body = v
	return s
}

func (s *CreateToolsetRequest) Validate() error {
	return dara.Validate(s)
}
