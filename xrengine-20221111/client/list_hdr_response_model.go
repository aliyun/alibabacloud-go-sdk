// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHdrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHdrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHdrResponse
	GetStatusCode() *int32
	SetBody(v *ListHdrResponseBody) *ListHdrResponse
	GetBody() *ListHdrResponseBody
}

type ListHdrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHdrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHdrResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHdrResponse) GoString() string {
	return s.String()
}

func (s *ListHdrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHdrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHdrResponse) GetBody() *ListHdrResponseBody {
	return s.Body
}

func (s *ListHdrResponse) SetHeaders(v map[string]*string) *ListHdrResponse {
	s.Headers = v
	return s
}

func (s *ListHdrResponse) SetStatusCode(v int32) *ListHdrResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHdrResponse) SetBody(v *ListHdrResponseBody) *ListHdrResponse {
	s.Body = v
	return s
}

func (s *ListHdrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
