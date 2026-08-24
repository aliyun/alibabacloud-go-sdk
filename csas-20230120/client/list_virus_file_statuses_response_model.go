// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusFileStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusFileStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusFileStatusesResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusFileStatusesResponseBody) *ListVirusFileStatusesResponse
	GetBody() *ListVirusFileStatusesResponseBody
}

type ListVirusFileStatusesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusFileStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusFileStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusFileStatusesResponse) GoString() string {
	return s.String()
}

func (s *ListVirusFileStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusFileStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusFileStatusesResponse) GetBody() *ListVirusFileStatusesResponseBody {
	return s.Body
}

func (s *ListVirusFileStatusesResponse) SetHeaders(v map[string]*string) *ListVirusFileStatusesResponse {
	s.Headers = v
	return s
}

func (s *ListVirusFileStatusesResponse) SetStatusCode(v int32) *ListVirusFileStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusFileStatusesResponse) SetBody(v *ListVirusFileStatusesResponseBody) *ListVirusFileStatusesResponse {
	s.Body = v
	return s
}

func (s *ListVirusFileStatusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
