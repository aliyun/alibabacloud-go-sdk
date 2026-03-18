// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreatePackageRequest
	GetBody() *string
	SetIsInstall(v bool) *CreatePackageRequest
	GetIsInstall() *bool
}

type CreatePackageRequest struct {
	Body      *string `json:"body,omitempty" xml:"body,omitempty"`
	IsInstall *bool   `json:"isInstall,omitempty" xml:"isInstall,omitempty"`
}

func (s CreatePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePackageRequest) GoString() string {
	return s.String()
}

func (s *CreatePackageRequest) GetBody() *string {
	return s.Body
}

func (s *CreatePackageRequest) GetIsInstall() *bool {
	return s.IsInstall
}

func (s *CreatePackageRequest) SetBody(v string) *CreatePackageRequest {
	s.Body = &v
	return s
}

func (s *CreatePackageRequest) SetIsInstall(v bool) *CreatePackageRequest {
	s.IsInstall = &v
	return s
}

func (s *CreatePackageRequest) Validate() error {
	return dara.Validate(s)
}
