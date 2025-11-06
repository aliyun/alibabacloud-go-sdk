// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcdpAimResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcdpAimResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcdpAimResponse
	GetStatusCode() *int32
	SetBody(v *ListMcdpAimResponseBody) *ListMcdpAimResponse
	GetBody() *ListMcdpAimResponseBody
}

type ListMcdpAimResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcdpAimResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcdpAimResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcdpAimResponse) GoString() string {
	return s.String()
}

func (s *ListMcdpAimResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcdpAimResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcdpAimResponse) GetBody() *ListMcdpAimResponseBody {
	return s.Body
}

func (s *ListMcdpAimResponse) SetHeaders(v map[string]*string) *ListMcdpAimResponse {
	s.Headers = v
	return s
}

func (s *ListMcdpAimResponse) SetStatusCode(v int32) *ListMcdpAimResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcdpAimResponse) SetBody(v *ListMcdpAimResponseBody) *ListMcdpAimResponse {
	s.Body = v
	return s
}

func (s *ListMcdpAimResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
