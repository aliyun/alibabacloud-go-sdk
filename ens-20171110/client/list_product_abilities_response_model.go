// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductAbilitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductAbilitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductAbilitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListProductAbilitiesResponseBody) *ListProductAbilitiesResponse
	GetBody() *ListProductAbilitiesResponseBody
}

type ListProductAbilitiesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductAbilitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductAbilitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductAbilitiesResponse) GoString() string {
	return s.String()
}

func (s *ListProductAbilitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductAbilitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductAbilitiesResponse) GetBody() *ListProductAbilitiesResponseBody {
	return s.Body
}

func (s *ListProductAbilitiesResponse) SetHeaders(v map[string]*string) *ListProductAbilitiesResponse {
	s.Headers = v
	return s
}

func (s *ListProductAbilitiesResponse) SetStatusCode(v int32) *ListProductAbilitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductAbilitiesResponse) SetBody(v *ListProductAbilitiesResponseBody) *ListProductAbilitiesResponse {
	s.Body = v
	return s
}

func (s *ListProductAbilitiesResponse) Validate() error {
	return dara.Validate(s)
}
