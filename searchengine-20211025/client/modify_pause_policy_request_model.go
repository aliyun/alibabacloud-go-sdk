// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPausePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]*BodyValue) *ModifyPausePolicyRequest
	GetBody() map[string]*BodyValue
}

type ModifyPausePolicyRequest struct {
	// The request body.
	Body map[string]*BodyValue `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPausePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPausePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPausePolicyRequest) GetBody() map[string]*BodyValue {
	return s.Body
}

func (s *ModifyPausePolicyRequest) SetBody(v map[string]*BodyValue) *ModifyPausePolicyRequest {
	s.Body = v
	return s
}

func (s *ModifyPausePolicyRequest) Validate() error {
	return dara.Validate(s)
}
