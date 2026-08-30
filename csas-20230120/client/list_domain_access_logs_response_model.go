// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainAccessLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDomainAccessLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDomainAccessLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListDomainAccessLogsResponseBody) *ListDomainAccessLogsResponse
	GetBody() *ListDomainAccessLogsResponseBody
}

type ListDomainAccessLogsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainAccessLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainAccessLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDomainAccessLogsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainAccessLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDomainAccessLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDomainAccessLogsResponse) GetBody() *ListDomainAccessLogsResponseBody {
	return s.Body
}

func (s *ListDomainAccessLogsResponse) SetHeaders(v map[string]*string) *ListDomainAccessLogsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainAccessLogsResponse) SetStatusCode(v int32) *ListDomainAccessLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainAccessLogsResponse) SetBody(v *ListDomainAccessLogsResponseBody) *ListDomainAccessLogsResponse {
	s.Body = v
	return s
}

func (s *ListDomainAccessLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
