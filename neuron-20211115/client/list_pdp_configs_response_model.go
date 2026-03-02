// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpConfigsResponse
	GetStatusCode() *int32
	SetBody(v *PdpConfigPageResult) *ListPdpConfigsResponse
	GetBody() *PdpConfigPageResult
}

type ListPdpConfigsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpConfigPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListPdpConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpConfigsResponse) GetBody() *PdpConfigPageResult {
	return s.Body
}

func (s *ListPdpConfigsResponse) SetHeaders(v map[string]*string) *ListPdpConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListPdpConfigsResponse) SetStatusCode(v int32) *ListPdpConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpConfigsResponse) SetBody(v *PdpConfigPageResult) *ListPdpConfigsResponse {
	s.Body = v
	return s
}

func (s *ListPdpConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
