// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableEsInstanceIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableEsInstanceIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableEsInstanceIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableEsInstanceIdsResponseBody) *ListAvailableEsInstanceIdsResponse
	GetBody() *ListAvailableEsInstanceIdsResponseBody
}

type ListAvailableEsInstanceIdsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableEsInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableEsInstanceIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableEsInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableEsInstanceIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableEsInstanceIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableEsInstanceIdsResponse) GetBody() *ListAvailableEsInstanceIdsResponseBody {
	return s.Body
}

func (s *ListAvailableEsInstanceIdsResponse) SetHeaders(v map[string]*string) *ListAvailableEsInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableEsInstanceIdsResponse) SetStatusCode(v int32) *ListAvailableEsInstanceIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponse) SetBody(v *ListAvailableEsInstanceIdsResponseBody) *ListAvailableEsInstanceIdsResponse {
	s.Body = v
	return s
}

func (s *ListAvailableEsInstanceIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
