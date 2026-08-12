// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRiskItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRiskItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListRiskItemsResponseBody) *ListRiskItemsResponse
	GetBody() *ListRiskItemsResponseBody
}

type ListRiskItemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRiskItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRiskItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRiskItemsResponse) GoString() string {
	return s.String()
}

func (s *ListRiskItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRiskItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRiskItemsResponse) GetBody() *ListRiskItemsResponseBody {
	return s.Body
}

func (s *ListRiskItemsResponse) SetHeaders(v map[string]*string) *ListRiskItemsResponse {
	s.Headers = v
	return s
}

func (s *ListRiskItemsResponse) SetStatusCode(v int32) *ListRiskItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRiskItemsResponse) SetBody(v *ListRiskItemsResponseBody) *ListRiskItemsResponse {
	s.Body = v
	return s
}

func (s *ListRiskItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
