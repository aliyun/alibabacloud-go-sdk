// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateKnowLedgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*string) *ValidateKnowLedgeRequest
	GetBody() []*string
}

type ValidateKnowLedgeRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ValidateKnowLedgeRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateKnowLedgeRequest) GoString() string {
	return s.String()
}

func (s *ValidateKnowLedgeRequest) GetBody() []*string {
	return s.Body
}

func (s *ValidateKnowLedgeRequest) SetBody(v []*string) *ValidateKnowLedgeRequest {
	s.Body = v
	return s
}

func (s *ValidateKnowLedgeRequest) Validate() error {
	return dara.Validate(s)
}
