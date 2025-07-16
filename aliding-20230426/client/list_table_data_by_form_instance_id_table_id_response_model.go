// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDataByFormInstanceIdTableIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableDataByFormInstanceIdTableIdResponse
	GetStatusCode() *int32
	SetBody(v *ListTableDataByFormInstanceIdTableIdResponseBody) *ListTableDataByFormInstanceIdTableIdResponse
	GetBody() *ListTableDataByFormInstanceIdTableIdResponseBody
}

type ListTableDataByFormInstanceIdTableIdResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableDataByFormInstanceIdTableIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdResponse) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) GetBody() *ListTableDataByFormInstanceIdTableIdResponseBody {
	return s.Body
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) SetHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdResponse {
	s.Headers = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) SetStatusCode(v int32) *ListTableDataByFormInstanceIdTableIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) SetBody(v *ListTableDataByFormInstanceIdTableIdResponseBody) *ListTableDataByFormInstanceIdTableIdResponse {
	s.Body = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) Validate() error {
	return dara.Validate(s)
}
