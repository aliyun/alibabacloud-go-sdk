// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportNameListResponse
	GetStatusCode() *int32
	SetBody(v *ImportNameListResponseBody) *ImportNameListResponse
	GetBody() *ImportNameListResponseBody
}

type ImportNameListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportNameListResponse) GoString() string {
	return s.String()
}

func (s *ImportNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportNameListResponse) GetBody() *ImportNameListResponseBody {
	return s.Body
}

func (s *ImportNameListResponse) SetHeaders(v map[string]*string) *ImportNameListResponse {
	s.Headers = v
	return s
}

func (s *ImportNameListResponse) SetStatusCode(v int32) *ImportNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportNameListResponse) SetBody(v *ImportNameListResponseBody) *ImportNameListResponse {
	s.Body = v
	return s
}

func (s *ImportNameListResponse) Validate() error {
	return dara.Validate(s)
}
