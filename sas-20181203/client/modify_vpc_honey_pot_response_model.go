// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcHoneyPotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcHoneyPotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcHoneyPotResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcHoneyPotResponseBody) *ModifyVpcHoneyPotResponse
	GetBody() *ModifyVpcHoneyPotResponseBody
}

type ModifyVpcHoneyPotResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcHoneyPotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcHoneyPotResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcHoneyPotResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcHoneyPotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcHoneyPotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcHoneyPotResponse) GetBody() *ModifyVpcHoneyPotResponseBody {
	return s.Body
}

func (s *ModifyVpcHoneyPotResponse) SetHeaders(v map[string]*string) *ModifyVpcHoneyPotResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcHoneyPotResponse) SetStatusCode(v int32) *ModifyVpcHoneyPotResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcHoneyPotResponse) SetBody(v *ModifyVpcHoneyPotResponseBody) *ModifyVpcHoneyPotResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcHoneyPotResponse) Validate() error {
	return dara.Validate(s)
}
