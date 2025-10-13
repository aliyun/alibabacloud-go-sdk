// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishedServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishedServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishedServicesResponseBody) *ListPublishedServicesResponse
	GetBody() *ListPublishedServicesResponseBody
}

type ListPublishedServicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishedServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishedServicesResponse) GetBody() *ListPublishedServicesResponseBody {
	return s.Body
}

func (s *ListPublishedServicesResponse) SetHeaders(v map[string]*string) *ListPublishedServicesResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedServicesResponse) SetStatusCode(v int32) *ListPublishedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedServicesResponse) SetBody(v *ListPublishedServicesResponseBody) *ListPublishedServicesResponse {
	s.Body = v
	return s
}

func (s *ListPublishedServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
