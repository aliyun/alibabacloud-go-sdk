// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcInvokesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcInvokesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcInvokesResponse
	GetStatusCode() *int32
	SetBody(v *ListPbcInvokesResponseBody) *ListPbcInvokesResponse
	GetBody() *ListPbcInvokesResponseBody
}

type ListPbcInvokesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPbcInvokesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcInvokesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcInvokesResponse) GoString() string {
	return s.String()
}

func (s *ListPbcInvokesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcInvokesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcInvokesResponse) GetBody() *ListPbcInvokesResponseBody {
	return s.Body
}

func (s *ListPbcInvokesResponse) SetHeaders(v map[string]*string) *ListPbcInvokesResponse {
	s.Headers = v
	return s
}

func (s *ListPbcInvokesResponse) SetStatusCode(v int32) *ListPbcInvokesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcInvokesResponse) SetBody(v *ListPbcInvokesResponseBody) *ListPbcInvokesResponse {
	s.Body = v
	return s
}

func (s *ListPbcInvokesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
