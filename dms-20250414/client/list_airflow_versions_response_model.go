// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAirflowVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAirflowVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAirflowVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAirflowVersionsResponseBody) *ListAirflowVersionsResponse
	GetBody() *ListAirflowVersionsResponseBody
}

type ListAirflowVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAirflowVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAirflowVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAirflowVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAirflowVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAirflowVersionsResponse) GetBody() *ListAirflowVersionsResponseBody {
	return s.Body
}

func (s *ListAirflowVersionsResponse) SetHeaders(v map[string]*string) *ListAirflowVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAirflowVersionsResponse) SetStatusCode(v int32) *ListAirflowVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAirflowVersionsResponse) SetBody(v *ListAirflowVersionsResponseBody) *ListAirflowVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAirflowVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
