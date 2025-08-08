// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceDeploymentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceDeploymentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceDeploymentsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceDeploymentsResponseBody) *ListServiceDeploymentsResponse
	GetBody() *ListServiceDeploymentsResponseBody
}

type ListServiceDeploymentsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceDeploymentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceDeploymentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceDeploymentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceDeploymentsResponse) GetBody() *ListServiceDeploymentsResponseBody {
	return s.Body
}

func (s *ListServiceDeploymentsResponse) SetHeaders(v map[string]*string) *ListServiceDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceDeploymentsResponse) SetStatusCode(v int32) *ListServiceDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceDeploymentsResponse) SetBody(v *ListServiceDeploymentsResponseBody) *ListServiceDeploymentsResponse {
	s.Body = v
	return s
}

func (s *ListServiceDeploymentsResponse) Validate() error {
	return dara.Validate(s)
}
