// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotViewPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotViewPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotViewPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListHotViewPointsResponseBody) *ListHotViewPointsResponse
	GetBody() *ListHotViewPointsResponseBody
}

type ListHotViewPointsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotViewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotViewPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotViewPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotViewPointsResponse) GetBody() *ListHotViewPointsResponseBody {
	return s.Body
}

func (s *ListHotViewPointsResponse) SetHeaders(v map[string]*string) *ListHotViewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListHotViewPointsResponse) SetStatusCode(v int32) *ListHotViewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotViewPointsResponse) SetBody(v *ListHotViewPointsResponseBody) *ListHotViewPointsResponse {
	s.Body = v
	return s
}

func (s *ListHotViewPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
