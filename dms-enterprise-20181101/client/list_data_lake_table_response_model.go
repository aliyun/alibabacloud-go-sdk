// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeTableResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeTableResponseBody) *ListDataLakeTableResponse
	GetBody() *ListDataLakeTableResponseBody
}

type ListDataLakeTableResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTableResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeTableResponse) GetBody() *ListDataLakeTableResponseBody {
	return s.Body
}

func (s *ListDataLakeTableResponse) SetHeaders(v map[string]*string) *ListDataLakeTableResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeTableResponse) SetStatusCode(v int32) *ListDataLakeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeTableResponse) SetBody(v *ListDataLakeTableResponseBody) *ListDataLakeTableResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
