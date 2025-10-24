// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsPartitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsPartitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsPartitionsResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsPartitionsResponseBody) *ListMmsPartitionsResponse
	GetBody() *ListMmsPartitionsResponseBody
}

type ListMmsPartitionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsPartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsPartitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsPartitionsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsPartitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsPartitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsPartitionsResponse) GetBody() *ListMmsPartitionsResponseBody {
	return s.Body
}

func (s *ListMmsPartitionsResponse) SetHeaders(v map[string]*string) *ListMmsPartitionsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsPartitionsResponse) SetStatusCode(v int32) *ListMmsPartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsPartitionsResponse) SetBody(v *ListMmsPartitionsResponseBody) *ListMmsPartitionsResponse {
	s.Body = v
	return s
}

func (s *ListMmsPartitionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
