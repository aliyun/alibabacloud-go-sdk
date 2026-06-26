// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *GetLoginPreferenceRequest
	GetUserPoolName() *string
}

type GetLoginPreferenceRequest struct {
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetLoginPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceRequest) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetLoginPreferenceRequest) SetUserPoolName(v string) *GetLoginPreferenceRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetLoginPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
