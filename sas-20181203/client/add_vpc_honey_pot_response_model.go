// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVpcHoneyPotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVpcHoneyPotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVpcHoneyPotResponse
	GetStatusCode() *int32
	SetBody(v *AddVpcHoneyPotResponseBody) *AddVpcHoneyPotResponse
	GetBody() *AddVpcHoneyPotResponseBody
}

type AddVpcHoneyPotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVpcHoneyPotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVpcHoneyPotResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVpcHoneyPotResponse) GoString() string {
	return s.String()
}

func (s *AddVpcHoneyPotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVpcHoneyPotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVpcHoneyPotResponse) GetBody() *AddVpcHoneyPotResponseBody {
	return s.Body
}

func (s *AddVpcHoneyPotResponse) SetHeaders(v map[string]*string) *AddVpcHoneyPotResponse {
	s.Headers = v
	return s
}

func (s *AddVpcHoneyPotResponse) SetStatusCode(v int32) *AddVpcHoneyPotResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVpcHoneyPotResponse) SetBody(v *AddVpcHoneyPotResponseBody) *AddVpcHoneyPotResponse {
	s.Body = v
	return s
}

func (s *AddVpcHoneyPotResponse) Validate() error {
	return dara.Validate(s)
}
