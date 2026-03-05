// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppTokenServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAction(v string) *CreateAppTokenServiceRequest
	GetCreateAction() *string
}

type CreateAppTokenServiceRequest struct {
	// example:
	//
	// OPEN_SUBSCRIPTION
	CreateAction *string `json:"CreateAction,omitempty" xml:"CreateAction,omitempty"`
}

func (s CreateAppTokenServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppTokenServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateAppTokenServiceRequest) GetCreateAction() *string {
	return s.CreateAction
}

func (s *CreateAppTokenServiceRequest) SetCreateAction(v string) *CreateAppTokenServiceRequest {
	s.CreateAction = &v
	return s
}

func (s *CreateAppTokenServiceRequest) Validate() error {
	return dara.Validate(s)
}
