// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPausePolicysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPausePolicysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPausePolicysResponse
	GetStatusCode() *int32
	SetBody(v *ListPausePolicysResponseBody) *ListPausePolicysResponse
	GetBody() *ListPausePolicysResponseBody
}

type ListPausePolicysResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPausePolicysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPausePolicysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPausePolicysResponse) GoString() string {
	return s.String()
}

func (s *ListPausePolicysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPausePolicysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPausePolicysResponse) GetBody() *ListPausePolicysResponseBody {
	return s.Body
}

func (s *ListPausePolicysResponse) SetHeaders(v map[string]*string) *ListPausePolicysResponse {
	s.Headers = v
	return s
}

func (s *ListPausePolicysResponse) SetStatusCode(v int32) *ListPausePolicysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPausePolicysResponse) SetBody(v *ListPausePolicysResponseBody) *ListPausePolicysResponse {
	s.Body = v
	return s
}

func (s *ListPausePolicysResponse) Validate() error {
	return dara.Validate(s)
}
