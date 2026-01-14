// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhitePortalMenuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWhitePortalMenuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWhitePortalMenuResponse
	GetStatusCode() *int32
	SetBody(v *ListWhitePortalMenuResponseBody) *ListWhitePortalMenuResponse
	GetBody() *ListWhitePortalMenuResponseBody
}

type ListWhitePortalMenuResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWhitePortalMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWhitePortalMenuResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWhitePortalMenuResponse) GoString() string {
	return s.String()
}

func (s *ListWhitePortalMenuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWhitePortalMenuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWhitePortalMenuResponse) GetBody() *ListWhitePortalMenuResponseBody {
	return s.Body
}

func (s *ListWhitePortalMenuResponse) SetHeaders(v map[string]*string) *ListWhitePortalMenuResponse {
	s.Headers = v
	return s
}

func (s *ListWhitePortalMenuResponse) SetStatusCode(v int32) *ListWhitePortalMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWhitePortalMenuResponse) SetBody(v *ListWhitePortalMenuResponseBody) *ListWhitePortalMenuResponse {
	s.Body = v
	return s
}

func (s *ListWhitePortalMenuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
