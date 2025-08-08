// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolsetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Toolset) *UpdateToolsetRequest
	GetBody() *Toolset
}

type UpdateToolsetRequest struct {
	Body *Toolset `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateToolsetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolsetRequest) GoString() string {
	return s.String()
}

func (s *UpdateToolsetRequest) GetBody() *Toolset {
	return s.Body
}

func (s *UpdateToolsetRequest) SetBody(v *Toolset) *UpdateToolsetRequest {
	s.Body = v
	return s
}

func (s *UpdateToolsetRequest) Validate() error {
	return dara.Validate(s)
}
