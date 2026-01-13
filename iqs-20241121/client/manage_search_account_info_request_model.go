// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSearchAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *AccountInfoManageRequest) *ManageSearchAccountInfoRequest
	GetBody() *AccountInfoManageRequest
}

type ManageSearchAccountInfoRequest struct {
	Body *AccountInfoManageRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSearchAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageSearchAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *ManageSearchAccountInfoRequest) GetBody() *AccountInfoManageRequest {
	return s.Body
}

func (s *ManageSearchAccountInfoRequest) SetBody(v *AccountInfoManageRequest) *ManageSearchAccountInfoRequest {
	s.Body = v
	return s
}

func (s *ManageSearchAccountInfoRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
