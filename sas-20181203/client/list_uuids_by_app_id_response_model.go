// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUuidsByAppIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUuidsByAppIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUuidsByAppIdResponse
	GetStatusCode() *int32
	SetBody(v *ListUuidsByAppIdResponseBody) *ListUuidsByAppIdResponse
	GetBody() *ListUuidsByAppIdResponseBody
}

type ListUuidsByAppIdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUuidsByAppIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUuidsByAppIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByAppIdResponse) GoString() string {
	return s.String()
}

func (s *ListUuidsByAppIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUuidsByAppIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUuidsByAppIdResponse) GetBody() *ListUuidsByAppIdResponseBody {
	return s.Body
}

func (s *ListUuidsByAppIdResponse) SetHeaders(v map[string]*string) *ListUuidsByAppIdResponse {
	s.Headers = v
	return s
}

func (s *ListUuidsByAppIdResponse) SetStatusCode(v int32) *ListUuidsByAppIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUuidsByAppIdResponse) SetBody(v *ListUuidsByAppIdResponseBody) *ListUuidsByAppIdResponse {
	s.Body = v
	return s
}

func (s *ListUuidsByAppIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
