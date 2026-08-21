// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConnectorResponse
	GetStatusCode() *int32
	SetBody(v *CreateConnectorResponseBody) *CreateConnectorResponse
	GetBody() *CreateConnectorResponseBody
}

type CreateConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectorResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConnectorResponse) GetBody() *CreateConnectorResponseBody {
	return s.Body
}

func (s *CreateConnectorResponse) SetHeaders(v map[string]*string) *CreateConnectorResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectorResponse) SetStatusCode(v int32) *CreateConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnectorResponse) SetBody(v *CreateConnectorResponseBody) *CreateConnectorResponse {
	s.Body = v
	return s
}

func (s *CreateConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
