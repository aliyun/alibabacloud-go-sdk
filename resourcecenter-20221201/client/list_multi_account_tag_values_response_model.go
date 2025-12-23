// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagValuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiAccountTagValuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiAccountTagValuesResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiAccountTagValuesResponseBody) *ListMultiAccountTagValuesResponse
	GetBody() *ListMultiAccountTagValuesResponseBody
}

type ListMultiAccountTagValuesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountTagValuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiAccountTagValuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiAccountTagValuesResponse) GetBody() *ListMultiAccountTagValuesResponseBody {
	return s.Body
}

func (s *ListMultiAccountTagValuesResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetStatusCode(v int32) *ListMultiAccountTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetBody(v *ListMultiAccountTagValuesResponseBody) *ListMultiAccountTagValuesResponse {
	s.Body = v
	return s
}

func (s *ListMultiAccountTagValuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
