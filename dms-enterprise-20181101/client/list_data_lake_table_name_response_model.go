// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTableNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeTableNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeTableNameResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeTableNameResponseBody) *ListDataLakeTableNameResponse
	GetBody() *ListDataLakeTableNameResponseBody
}

type ListDataLakeTableNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeTableNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeTableNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTableNameResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeTableNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeTableNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeTableNameResponse) GetBody() *ListDataLakeTableNameResponseBody {
	return s.Body
}

func (s *ListDataLakeTableNameResponse) SetHeaders(v map[string]*string) *ListDataLakeTableNameResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeTableNameResponse) SetStatusCode(v int32) *ListDataLakeTableNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeTableNameResponse) SetBody(v *ListDataLakeTableNameResponseBody) *ListDataLakeTableNameResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeTableNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
