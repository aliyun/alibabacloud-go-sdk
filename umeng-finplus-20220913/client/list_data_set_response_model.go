// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSetResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSetResponseBody) *ListDataSetResponse
	GetBody() *ListDataSetResponseBody
}

type ListDataSetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponse) GoString() string {
	return s.String()
}

func (s *ListDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSetResponse) GetBody() *ListDataSetResponseBody {
	return s.Body
}

func (s *ListDataSetResponse) SetHeaders(v map[string]*string) *ListDataSetResponse {
	s.Headers = v
	return s
}

func (s *ListDataSetResponse) SetStatusCode(v int32) *ListDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSetResponse) SetBody(v *ListDataSetResponseBody) *ListDataSetResponse {
	s.Body = v
	return s
}

func (s *ListDataSetResponse) Validate() error {
	return dara.Validate(s)
}
