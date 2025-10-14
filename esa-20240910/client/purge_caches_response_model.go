// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurgeCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurgeCachesResponse
	GetStatusCode() *int32
	SetBody(v *PurgeCachesResponseBody) *PurgeCachesResponse
	GetBody() *PurgeCachesResponseBody
}

type PurgeCachesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurgeCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurgeCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesResponse) GoString() string {
	return s.String()
}

func (s *PurgeCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurgeCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurgeCachesResponse) GetBody() *PurgeCachesResponseBody {
	return s.Body
}

func (s *PurgeCachesResponse) SetHeaders(v map[string]*string) *PurgeCachesResponse {
	s.Headers = v
	return s
}

func (s *PurgeCachesResponse) SetStatusCode(v int32) *PurgeCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PurgeCachesResponse) SetBody(v *PurgeCachesResponseBody) *PurgeCachesResponse {
	s.Body = v
	return s
}

func (s *PurgeCachesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
