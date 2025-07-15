// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialoguesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDialoguesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDialoguesResponse
	GetStatusCode() *int32
	SetBody(v *ListDialoguesResponseBody) *ListDialoguesResponse
	GetBody() *ListDialoguesResponseBody
}

type ListDialoguesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialoguesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialoguesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDialoguesResponse) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDialoguesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDialoguesResponse) GetBody() *ListDialoguesResponseBody {
	return s.Body
}

func (s *ListDialoguesResponse) SetHeaders(v map[string]*string) *ListDialoguesResponse {
	s.Headers = v
	return s
}

func (s *ListDialoguesResponse) SetStatusCode(v int32) *ListDialoguesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialoguesResponse) SetBody(v *ListDialoguesResponseBody) *ListDialoguesResponse {
	s.Body = v
	return s
}

func (s *ListDialoguesResponse) Validate() error {
	return dara.Validate(s)
}
