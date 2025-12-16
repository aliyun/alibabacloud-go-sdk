// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTableFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceTableFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceTableFieldsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceTableFieldsResponseBody) *ListDataSourceTableFieldsResponse
	GetBody() *ListDataSourceTableFieldsResponseBody
}

type ListDataSourceTableFieldsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTableFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTableFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTableFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceTableFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceTableFieldsResponse) GetBody() *ListDataSourceTableFieldsResponseBody {
	return s.Body
}

func (s *ListDataSourceTableFieldsResponse) SetHeaders(v map[string]*string) *ListDataSourceTableFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTableFieldsResponse) SetStatusCode(v int32) *ListDataSourceTableFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTableFieldsResponse) SetBody(v *ListDataSourceTableFieldsResponseBody) *ListDataSourceTableFieldsResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceTableFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
