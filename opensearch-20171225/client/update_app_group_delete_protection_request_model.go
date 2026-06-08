// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppGroupDeleteProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *UpdateAppGroupDeleteProtectionRequest
	GetBody() map[string]interface{}
}

type UpdateAppGroupDeleteProtectionRequest struct {
	// example:
	//
	// {
	//
	//   "deleteProtection": "on"
	//
	// }
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppGroupDeleteProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupDeleteProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupDeleteProtectionRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *UpdateAppGroupDeleteProtectionRequest) SetBody(v map[string]interface{}) *UpdateAppGroupDeleteProtectionRequest {
	s.Body = v
	return s
}

func (s *UpdateAppGroupDeleteProtectionRequest) Validate() error {
	return dara.Validate(s)
}
