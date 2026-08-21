// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceInfoResponseBody) *ListInstanceInfoResponse
	GetBody() *ListInstanceInfoResponseBody
}

type ListInstanceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceInfoResponse) GetBody() *ListInstanceInfoResponseBody {
	return s.Body
}

func (s *ListInstanceInfoResponse) SetHeaders(v map[string]*string) *ListInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceInfoResponse) SetStatusCode(v int32) *ListInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceInfoResponse) SetBody(v *ListInstanceInfoResponseBody) *ListInstanceInfoResponse {
	s.Body = v
	return s
}

func (s *ListInstanceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
