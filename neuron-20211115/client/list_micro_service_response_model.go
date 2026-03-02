// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMicroServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMicroServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMicroServiceResponse
	GetStatusCode() *int32
	SetBody(v *MicroServicePageResult) *ListMicroServiceResponse
	GetBody() *MicroServicePageResult
}

type ListMicroServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MicroServicePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMicroServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMicroServiceResponse) GoString() string {
	return s.String()
}

func (s *ListMicroServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMicroServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMicroServiceResponse) GetBody() *MicroServicePageResult {
	return s.Body
}

func (s *ListMicroServiceResponse) SetHeaders(v map[string]*string) *ListMicroServiceResponse {
	s.Headers = v
	return s
}

func (s *ListMicroServiceResponse) SetStatusCode(v int32) *ListMicroServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMicroServiceResponse) SetBody(v *MicroServicePageResult) *ListMicroServiceResponse {
	s.Body = v
	return s
}

func (s *ListMicroServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
