// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTablebaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeTablebaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeTablebaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeTablebaseInfoResponseBody) *ListDataLakeTablebaseInfoResponse
	GetBody() *ListDataLakeTablebaseInfoResponseBody
}

type ListDataLakeTablebaseInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeTablebaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeTablebaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTablebaseInfoResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeTablebaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeTablebaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeTablebaseInfoResponse) GetBody() *ListDataLakeTablebaseInfoResponseBody {
	return s.Body
}

func (s *ListDataLakeTablebaseInfoResponse) SetHeaders(v map[string]*string) *ListDataLakeTablebaseInfoResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeTablebaseInfoResponse) SetStatusCode(v int32) *ListDataLakeTablebaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeTablebaseInfoResponse) SetBody(v *ListDataLakeTablebaseInfoResponseBody) *ListDataLakeTablebaseInfoResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeTablebaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
