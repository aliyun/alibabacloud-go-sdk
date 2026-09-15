// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyConnectorResponse
	GetStatusCode() *int32
	SetBody(v *VerifyConnectorResponseBody) *VerifyConnectorResponse
	GetBody() *VerifyConnectorResponseBody
}

type VerifyConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyConnectorResponse) GoString() string {
	return s.String()
}

func (s *VerifyConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyConnectorResponse) GetBody() *VerifyConnectorResponseBody {
	return s.Body
}

func (s *VerifyConnectorResponse) SetHeaders(v map[string]*string) *VerifyConnectorResponse {
	s.Headers = v
	return s
}

func (s *VerifyConnectorResponse) SetStatusCode(v int32) *VerifyConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyConnectorResponse) SetBody(v *VerifyConnectorResponseBody) *VerifyConnectorResponse {
	s.Body = v
	return s
}

func (s *VerifyConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
