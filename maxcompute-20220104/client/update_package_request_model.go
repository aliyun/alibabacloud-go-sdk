// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePackageRequest
	GetBody() *string
}

type UpdatePackageRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePackageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePackageRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePackageRequest) SetBody(v string) *UpdatePackageRequest {
	s.Body = &v
	return s
}

func (s *UpdatePackageRequest) Validate() error {
	return dara.Validate(s)
}
