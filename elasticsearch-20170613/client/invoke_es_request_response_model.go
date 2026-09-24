// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeEsRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeEsRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeEsRequestResponse
	GetStatusCode() *int32
	SetBody(v *InvokeEsRequestResponseBody) *InvokeEsRequestResponse
	GetBody() *InvokeEsRequestResponseBody
}

type InvokeEsRequestResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeEsRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeEsRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeEsRequestResponse) GoString() string {
	return s.String()
}

func (s *InvokeEsRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeEsRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeEsRequestResponse) GetBody() *InvokeEsRequestResponseBody {
	return s.Body
}

func (s *InvokeEsRequestResponse) SetHeaders(v map[string]*string) *InvokeEsRequestResponse {
	s.Headers = v
	return s
}

func (s *InvokeEsRequestResponse) SetStatusCode(v int32) *InvokeEsRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeEsRequestResponse) SetBody(v *InvokeEsRequestResponseBody) *InvokeEsRequestResponse {
	s.Body = v
	return s
}

func (s *InvokeEsRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
