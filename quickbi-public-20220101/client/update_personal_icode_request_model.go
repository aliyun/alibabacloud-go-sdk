// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePersonalICodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuickBIUserId(v string) *UpdatePersonalICodeRequest
	GetQuickBIUserId() *string
}

type UpdatePersonalICodeRequest struct {
	// The user ID in Quick BI.
	//
	// example:
	//
	// sasdasasd
	QuickBIUserId *string `json:"QuickBIUserId,omitempty" xml:"QuickBIUserId,omitempty"`
}

func (s UpdatePersonalICodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePersonalICodeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePersonalICodeRequest) GetQuickBIUserId() *string {
	return s.QuickBIUserId
}

func (s *UpdatePersonalICodeRequest) SetQuickBIUserId(v string) *UpdatePersonalICodeRequest {
	s.QuickBIUserId = &v
	return s
}

func (s *UpdatePersonalICodeRequest) Validate() error {
	return dara.Validate(s)
}
