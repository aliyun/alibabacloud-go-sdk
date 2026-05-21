// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConnectorResponse
	GetStatusCode() *int32
	SetBody(v *GetConnectorResponseBody) *GetConnectorResponse
	GetBody() *GetConnectorResponseBody
}

type GetConnectorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponse) GoString() string {
	return s.String()
}

func (s *GetConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConnectorResponse) GetBody() *GetConnectorResponseBody {
	return s.Body
}

func (s *GetConnectorResponse) SetHeaders(v map[string]*string) *GetConnectorResponse {
	s.Headers = v
	return s
}

func (s *GetConnectorResponse) SetStatusCode(v int32) *GetConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectorResponse) SetBody(v *GetConnectorResponseBody) *GetConnectorResponse {
	s.Body = v
	return s
}

func (s *GetConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
