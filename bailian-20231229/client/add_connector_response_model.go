// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddConnectorResponse
	GetStatusCode() *int32
	SetBody(v *AddConnectorResponseBody) *AddConnectorResponse
	GetBody() *AddConnectorResponseBody
}

type AddConnectorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s AddConnectorResponse) GoString() string {
	return s.String()
}

func (s *AddConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddConnectorResponse) GetBody() *AddConnectorResponseBody {
	return s.Body
}

func (s *AddConnectorResponse) SetHeaders(v map[string]*string) *AddConnectorResponse {
	s.Headers = v
	return s
}

func (s *AddConnectorResponse) SetStatusCode(v int32) *AddConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *AddConnectorResponse) SetBody(v *AddConnectorResponseBody) *AddConnectorResponse {
	s.Body = v
	return s
}

func (s *AddConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
