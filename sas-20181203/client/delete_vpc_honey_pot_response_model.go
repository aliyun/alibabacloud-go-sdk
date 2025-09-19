// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcHoneyPotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcHoneyPotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcHoneyPotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcHoneyPotResponseBody) *DeleteVpcHoneyPotResponse
	GetBody() *DeleteVpcHoneyPotResponseBody
}

type DeleteVpcHoneyPotResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcHoneyPotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcHoneyPotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcHoneyPotResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcHoneyPotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcHoneyPotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcHoneyPotResponse) GetBody() *DeleteVpcHoneyPotResponseBody {
	return s.Body
}

func (s *DeleteVpcHoneyPotResponse) SetHeaders(v map[string]*string) *DeleteVpcHoneyPotResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcHoneyPotResponse) SetStatusCode(v int32) *DeleteVpcHoneyPotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcHoneyPotResponse) SetBody(v *DeleteVpcHoneyPotResponseBody) *DeleteVpcHoneyPotResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcHoneyPotResponse) Validate() error {
	return dara.Validate(s)
}
