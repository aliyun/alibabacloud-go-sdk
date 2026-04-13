// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDeploymentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackDeploymentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackDeploymentsResponse
	GetStatusCode() *int32
	SetBody(v *GetStackDeploymentsResponseBody) *GetStackDeploymentsResponse
	GetBody() *GetStackDeploymentsResponseBody
}

type GetStackDeploymentsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackDeploymentsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackDeploymentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackDeploymentsResponse) GetBody() *GetStackDeploymentsResponseBody {
	return s.Body
}

func (s *GetStackDeploymentsResponse) SetHeaders(v map[string]*string) *GetStackDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *GetStackDeploymentsResponse) SetStatusCode(v int32) *GetStackDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackDeploymentsResponse) SetBody(v *GetStackDeploymentsResponseBody) *GetStackDeploymentsResponse {
	s.Body = v
	return s
}

func (s *GetStackDeploymentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
