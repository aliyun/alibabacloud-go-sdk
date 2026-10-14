// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticTableBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticTableBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticTableBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticTableBaseInfoResponseBody) *ListAgenticTableBaseInfoResponse
	GetBody() *ListAgenticTableBaseInfoResponseBody
}

type ListAgenticTableBaseInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticTableBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticTableBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticTableBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticTableBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticTableBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticTableBaseInfoResponse) GetBody() *ListAgenticTableBaseInfoResponseBody {
	return s.Body
}

func (s *ListAgenticTableBaseInfoResponse) SetHeaders(v map[string]*string) *ListAgenticTableBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticTableBaseInfoResponse) SetStatusCode(v int32) *ListAgenticTableBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponse) SetBody(v *ListAgenticTableBaseInfoResponseBody) *ListAgenticTableBaseInfoResponse {
	s.Body = v
	return s
}

func (s *ListAgenticTableBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
