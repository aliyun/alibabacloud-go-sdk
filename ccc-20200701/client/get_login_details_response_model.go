// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoginDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoginDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetLoginDetailsResponseBody) *GetLoginDetailsResponse
	GetBody() *GetLoginDetailsResponseBody
}

type GetLoginDetailsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoginDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoginDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoginDetailsResponse) GetBody() *GetLoginDetailsResponseBody {
	return s.Body
}

func (s *GetLoginDetailsResponse) SetHeaders(v map[string]*string) *GetLoginDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetLoginDetailsResponse) SetStatusCode(v int32) *GetLoginDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginDetailsResponse) SetBody(v *GetLoginDetailsResponseBody) *GetLoginDetailsResponse {
	s.Body = v
	return s
}

func (s *GetLoginDetailsResponse) Validate() error {
	return dara.Validate(s)
}
