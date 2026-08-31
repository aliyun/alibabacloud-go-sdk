// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryBillingDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryBillingDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryBillingDetailsResponseBody) *ModelRouterQueryBillingDetailsResponse
	GetBody() *ModelRouterQueryBillingDetailsResponseBody
}

type ModelRouterQueryBillingDetailsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryBillingDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryBillingDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingDetailsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryBillingDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryBillingDetailsResponse) GetBody() *ModelRouterQueryBillingDetailsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryBillingDetailsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryBillingDetailsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponse) SetStatusCode(v int32) *ModelRouterQueryBillingDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponse) SetBody(v *ModelRouterQueryBillingDetailsResponseBody) *ModelRouterQueryBillingDetailsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryBillingDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
