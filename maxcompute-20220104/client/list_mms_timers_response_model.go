// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTimersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsTimersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsTimersResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsTimersResponseBody) *ListMmsTimersResponse
	GetBody() *ListMmsTimersResponseBody
}

type ListMmsTimersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTimersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTimersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimersResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTimersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsTimersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsTimersResponse) GetBody() *ListMmsTimersResponseBody {
	return s.Body
}

func (s *ListMmsTimersResponse) SetHeaders(v map[string]*string) *ListMmsTimersResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTimersResponse) SetStatusCode(v int32) *ListMmsTimersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTimersResponse) SetBody(v *ListMmsTimersResponseBody) *ListMmsTimersResponse {
	s.Body = v
	return s
}

func (s *ListMmsTimersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
