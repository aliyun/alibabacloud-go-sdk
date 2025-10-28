// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelVersionsResponseBody) *ListModelVersionsResponse
	GetBody() *ListModelVersionsResponseBody
}

type ListModelVersionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListModelVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelVersionsResponse) GetBody() *ListModelVersionsResponseBody {
	return s.Body
}

func (s *ListModelVersionsResponse) SetHeaders(v map[string]*string) *ListModelVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListModelVersionsResponse) SetStatusCode(v int32) *ListModelVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelVersionsResponse) SetBody(v *ListModelVersionsResponseBody) *ListModelVersionsResponse {
	s.Body = v
	return s
}

func (s *ListModelVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
