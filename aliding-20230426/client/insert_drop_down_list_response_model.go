// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDropDownListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertDropDownListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertDropDownListResponse
	GetStatusCode() *int32
	SetBody(v *InsertDropDownListResponseBody) *InsertDropDownListResponse
	GetBody() *InsertDropDownListResponseBody
}

type InsertDropDownListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertDropDownListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertDropDownListResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListResponse) GoString() string {
	return s.String()
}

func (s *InsertDropDownListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertDropDownListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertDropDownListResponse) GetBody() *InsertDropDownListResponseBody {
	return s.Body
}

func (s *InsertDropDownListResponse) SetHeaders(v map[string]*string) *InsertDropDownListResponse {
	s.Headers = v
	return s
}

func (s *InsertDropDownListResponse) SetStatusCode(v int32) *InsertDropDownListResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertDropDownListResponse) SetBody(v *InsertDropDownListResponseBody) *InsertDropDownListResponse {
	s.Body = v
	return s
}

func (s *InsertDropDownListResponse) Validate() error {
	return dara.Validate(s)
}
