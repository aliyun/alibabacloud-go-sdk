// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableLevelResponse
	GetStatusCode() *int32
	SetBody(v *ListTableLevelResponseBody) *ListTableLevelResponse
	GetBody() *ListTableLevelResponseBody
}

type ListTableLevelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableLevelResponse) GoString() string {
	return s.String()
}

func (s *ListTableLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableLevelResponse) GetBody() *ListTableLevelResponseBody {
	return s.Body
}

func (s *ListTableLevelResponse) SetHeaders(v map[string]*string) *ListTableLevelResponse {
	s.Headers = v
	return s
}

func (s *ListTableLevelResponse) SetStatusCode(v int32) *ListTableLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableLevelResponse) SetBody(v *ListTableLevelResponseBody) *ListTableLevelResponse {
	s.Body = v
	return s
}

func (s *ListTableLevelResponse) Validate() error {
	return dara.Validate(s)
}
