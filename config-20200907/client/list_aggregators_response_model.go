// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregatorsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregatorsResponseBody) *ListAggregatorsResponse
	GetBody() *ListAggregatorsResponseBody
}

type ListAggregatorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregatorsResponse) GetBody() *ListAggregatorsResponseBody {
	return s.Body
}

func (s *ListAggregatorsResponse) SetHeaders(v map[string]*string) *ListAggregatorsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregatorsResponse) SetStatusCode(v int32) *ListAggregatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregatorsResponse) SetBody(v *ListAggregatorsResponseBody) *ListAggregatorsResponse {
	s.Body = v
	return s
}

func (s *ListAggregatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
