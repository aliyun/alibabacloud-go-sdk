// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUuidsByWebPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUuidsByWebPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUuidsByWebPathResponse
	GetStatusCode() *int32
	SetBody(v *ListUuidsByWebPathResponseBody) *ListUuidsByWebPathResponse
	GetBody() *ListUuidsByWebPathResponseBody
}

type ListUuidsByWebPathResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUuidsByWebPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUuidsByWebPathResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByWebPathResponse) GoString() string {
	return s.String()
}

func (s *ListUuidsByWebPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUuidsByWebPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUuidsByWebPathResponse) GetBody() *ListUuidsByWebPathResponseBody {
	return s.Body
}

func (s *ListUuidsByWebPathResponse) SetHeaders(v map[string]*string) *ListUuidsByWebPathResponse {
	s.Headers = v
	return s
}

func (s *ListUuidsByWebPathResponse) SetStatusCode(v int32) *ListUuidsByWebPathResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUuidsByWebPathResponse) SetBody(v *ListUuidsByWebPathResponseBody) *ListUuidsByWebPathResponse {
	s.Body = v
	return s
}

func (s *ListUuidsByWebPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
