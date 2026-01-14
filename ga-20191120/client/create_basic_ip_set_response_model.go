// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicIpSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicIpSetResponseBody) *CreateBasicIpSetResponse
	GetBody() *CreateBasicIpSetResponseBody
}

type CreateBasicIpSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicIpSetResponse) GetBody() *CreateBasicIpSetResponseBody {
	return s.Body
}

func (s *CreateBasicIpSetResponse) SetHeaders(v map[string]*string) *CreateBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicIpSetResponse) SetStatusCode(v int32) *CreateBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicIpSetResponse) SetBody(v *CreateBasicIpSetResponseBody) *CreateBasicIpSetResponse {
	s.Body = v
	return s
}

func (s *CreateBasicIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
