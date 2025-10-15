// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishedAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishedAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishedAppInfoResponseBody) *ListPublishedAppInfoResponse
	GetBody() *ListPublishedAppInfoResponseBody
}

type ListPublishedAppInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishedAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishedAppInfoResponse) GetBody() *ListPublishedAppInfoResponseBody {
	return s.Body
}

func (s *ListPublishedAppInfoResponse) SetHeaders(v map[string]*string) *ListPublishedAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAppInfoResponse) SetStatusCode(v int32) *ListPublishedAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAppInfoResponse) SetBody(v *ListPublishedAppInfoResponseBody) *ListPublishedAppInfoResponse {
	s.Body = v
	return s
}

func (s *ListPublishedAppInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
