// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactName(v string) *DeleteContactRequest
	GetContactName() *string
	SetDryRun(v bool) *DeleteContactRequest
	GetDryRun() *bool
}

type DeleteContactRequest struct {
	// The name of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// John Doe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Specifies whether to perform only a dry run for this request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *DeleteContactRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteContactRequest) SetContactName(v string) *DeleteContactRequest {
	s.ContactName = &v
	return s
}

func (s *DeleteContactRequest) SetDryRun(v bool) *DeleteContactRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteContactRequest) Validate() error {
	return dara.Validate(s)
}
