// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpretersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCodeInterpretersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCodeInterpretersResponse
	GetStatusCode() *int32
	SetBody(v *ListCodeInterpretersResult) *ListCodeInterpretersResponse
	GetBody() *ListCodeInterpretersResult
}

type ListCodeInterpretersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCodeInterpretersResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCodeInterpretersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpretersResponse) GoString() string {
	return s.String()
}

func (s *ListCodeInterpretersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCodeInterpretersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCodeInterpretersResponse) GetBody() *ListCodeInterpretersResult {
	return s.Body
}

func (s *ListCodeInterpretersResponse) SetHeaders(v map[string]*string) *ListCodeInterpretersResponse {
	s.Headers = v
	return s
}

func (s *ListCodeInterpretersResponse) SetStatusCode(v int32) *ListCodeInterpretersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCodeInterpretersResponse) SetBody(v *ListCodeInterpretersResult) *ListCodeInterpretersResponse {
	s.Body = v
	return s
}

func (s *ListCodeInterpretersResponse) Validate() error {
	return dara.Validate(s)
}
