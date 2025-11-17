// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAuthorizationMenuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAuthorizationMenuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAuthorizationMenuResponse
	GetStatusCode() *int32
	SetBody(v *CancelAuthorizationMenuResponseBody) *CancelAuthorizationMenuResponse
	GetBody() *CancelAuthorizationMenuResponseBody
}

type CancelAuthorizationMenuResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAuthorizationMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAuthorizationMenuResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAuthorizationMenuResponse) GoString() string {
	return s.String()
}

func (s *CancelAuthorizationMenuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAuthorizationMenuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAuthorizationMenuResponse) GetBody() *CancelAuthorizationMenuResponseBody {
	return s.Body
}

func (s *CancelAuthorizationMenuResponse) SetHeaders(v map[string]*string) *CancelAuthorizationMenuResponse {
	s.Headers = v
	return s
}

func (s *CancelAuthorizationMenuResponse) SetStatusCode(v int32) *CancelAuthorizationMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAuthorizationMenuResponse) SetBody(v *CancelAuthorizationMenuResponseBody) *CancelAuthorizationMenuResponse {
	s.Body = v
	return s
}

func (s *CancelAuthorizationMenuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
