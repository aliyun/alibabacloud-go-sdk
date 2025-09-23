// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListValueAddedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListValueAddedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListValueAddedResponse
	GetStatusCode() *int32
	SetBody(v *ListValueAddedResponseBody) *ListValueAddedResponse
	GetBody() *ListValueAddedResponseBody
}

type ListValueAddedResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListValueAddedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListValueAddedResponse) String() string {
	return dara.Prettify(s)
}

func (s ListValueAddedResponse) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListValueAddedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListValueAddedResponse) GetBody() *ListValueAddedResponseBody {
	return s.Body
}

func (s *ListValueAddedResponse) SetHeaders(v map[string]*string) *ListValueAddedResponse {
	s.Headers = v
	return s
}

func (s *ListValueAddedResponse) SetStatusCode(v int32) *ListValueAddedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListValueAddedResponse) SetBody(v *ListValueAddedResponseBody) *ListValueAddedResponse {
	s.Body = v
	return s
}

func (s *ListValueAddedResponse) Validate() error {
	return dara.Validate(s)
}
