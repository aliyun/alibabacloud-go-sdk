// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *PushDocumentsResponseBody) *PushDocumentsResponse
	GetBody() *PushDocumentsResponseBody
}

type PushDocumentsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s PushDocumentsResponse) GoString() string {
	return s.String()
}

func (s *PushDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushDocumentsResponse) GetBody() *PushDocumentsResponseBody {
	return s.Body
}

func (s *PushDocumentsResponse) SetHeaders(v map[string]*string) *PushDocumentsResponse {
	s.Headers = v
	return s
}

func (s *PushDocumentsResponse) SetStatusCode(v int32) *PushDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDocumentsResponse) SetBody(v *PushDocumentsResponseBody) *PushDocumentsResponse {
	s.Body = v
	return s
}

func (s *PushDocumentsResponse) Validate() error {
	return dara.Validate(s)
}
