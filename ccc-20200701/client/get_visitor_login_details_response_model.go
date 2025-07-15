// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVisitorLoginDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVisitorLoginDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVisitorLoginDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetVisitorLoginDetailsResponseBody) *GetVisitorLoginDetailsResponse
	GetBody() *GetVisitorLoginDetailsResponseBody
}

type GetVisitorLoginDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVisitorLoginDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVisitorLoginDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVisitorLoginDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetVisitorLoginDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVisitorLoginDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVisitorLoginDetailsResponse) GetBody() *GetVisitorLoginDetailsResponseBody {
	return s.Body
}

func (s *GetVisitorLoginDetailsResponse) SetHeaders(v map[string]*string) *GetVisitorLoginDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetVisitorLoginDetailsResponse) SetStatusCode(v int32) *GetVisitorLoginDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVisitorLoginDetailsResponse) SetBody(v *GetVisitorLoginDetailsResponseBody) *GetVisitorLoginDetailsResponse {
	s.Body = v
	return s
}

func (s *GetVisitorLoginDetailsResponse) Validate() error {
	return dara.Validate(s)
}
