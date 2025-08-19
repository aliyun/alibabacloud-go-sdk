// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateTriggerInput) *UpdateTriggerRequest
	GetBody() *UpdateTriggerInput
}

type UpdateTriggerRequest struct {
	// The trigger configurations.
	//
	// This parameter is required.
	Body *UpdateTriggerInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) GetBody() *UpdateTriggerInput {
	return s.Body
}

func (s *UpdateTriggerRequest) SetBody(v *UpdateTriggerInput) *UpdateTriggerRequest {
	s.Body = v
	return s
}

func (s *UpdateTriggerRequest) Validate() error {
	return dara.Validate(s)
}
