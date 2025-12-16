// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupIdentity(v string) *GetDomainRequest
	GetAppGroupIdentity() *string
}

type GetDomainRequest struct {
	// The name or ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_app_group_name
	AppGroupIdentity *string `json:"appGroupIdentity,omitempty" xml:"appGroupIdentity,omitempty"`
}

func (s GetDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) GetAppGroupIdentity() *string {
	return s.AppGroupIdentity
}

func (s *GetDomainRequest) SetAppGroupIdentity(v string) *GetDomainRequest {
	s.AppGroupIdentity = &v
	return s
}

func (s *GetDomainRequest) Validate() error {
	return dara.Validate(s)
}
