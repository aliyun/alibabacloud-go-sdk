// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrowdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrowdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrowdsResponse
	GetStatusCode() *int32
	SetBody(v *ListCrowdsResponseBody) *ListCrowdsResponse
	GetBody() *ListCrowdsResponseBody
}

type ListCrowdsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrowdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrowdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdsResponse) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrowdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrowdsResponse) GetBody() *ListCrowdsResponseBody {
	return s.Body
}

func (s *ListCrowdsResponse) SetHeaders(v map[string]*string) *ListCrowdsResponse {
	s.Headers = v
	return s
}

func (s *ListCrowdsResponse) SetStatusCode(v int32) *ListCrowdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrowdsResponse) SetBody(v *ListCrowdsResponseBody) *ListCrowdsResponse {
	s.Body = v
	return s
}

func (s *ListCrowdsResponse) Validate() error {
	return dara.Validate(s)
}
