// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateTriggerInput) *CreateTriggerRequest
	GetBody() *CreateTriggerInput
}

type CreateTriggerRequest struct {
	// The trigger configurations.
	//
	// This parameter is required.
	Body *CreateTriggerInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) GetBody() *CreateTriggerInput {
	return s.Body
}

func (s *CreateTriggerRequest) SetBody(v *CreateTriggerInput) *CreateTriggerRequest {
	s.Body = v
	return s
}

func (s *CreateTriggerRequest) Validate() error {
	return dara.Validate(s)
}
