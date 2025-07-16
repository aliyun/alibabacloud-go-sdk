// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLabelTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLabelTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListLabelTablesResponseBody) *ListLabelTablesResponse
	GetBody() *ListLabelTablesResponseBody
}

type ListLabelTablesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLabelTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLabelTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLabelTablesResponse) GoString() string {
	return s.String()
}

func (s *ListLabelTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLabelTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLabelTablesResponse) GetBody() *ListLabelTablesResponseBody {
	return s.Body
}

func (s *ListLabelTablesResponse) SetHeaders(v map[string]*string) *ListLabelTablesResponse {
	s.Headers = v
	return s
}

func (s *ListLabelTablesResponse) SetStatusCode(v int32) *ListLabelTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLabelTablesResponse) SetBody(v *ListLabelTablesResponseBody) *ListLabelTablesResponse {
	s.Body = v
	return s
}

func (s *ListLabelTablesResponse) Validate() error {
	return dara.Validate(s)
}
