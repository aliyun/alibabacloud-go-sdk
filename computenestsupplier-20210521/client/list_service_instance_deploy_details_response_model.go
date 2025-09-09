// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceDeployDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceDeployDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceDeployDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceDeployDetailsResponseBody) *ListServiceInstanceDeployDetailsResponse
	GetBody() *ListServiceInstanceDeployDetailsResponseBody
}

type ListServiceInstanceDeployDetailsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceDeployDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceDeployDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceDeployDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceDeployDetailsResponse) GetBody() *ListServiceInstanceDeployDetailsResponseBody {
	return s.Body
}

func (s *ListServiceInstanceDeployDetailsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceDeployDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponse) SetStatusCode(v int32) *ListServiceInstanceDeployDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponse) SetBody(v *ListServiceInstanceDeployDetailsResponseBody) *ListServiceInstanceDeployDetailsResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponse) Validate() error {
	return dara.Validate(s)
}
