// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAncestorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAncestorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAncestorsResponse
	GetStatusCode() *int32
	SetBody(v *ListAncestorsResponseBody) *ListAncestorsResponse
	GetBody() *ListAncestorsResponseBody
}

type ListAncestorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAncestorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAncestorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAncestorsResponse) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAncestorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAncestorsResponse) GetBody() *ListAncestorsResponseBody {
	return s.Body
}

func (s *ListAncestorsResponse) SetHeaders(v map[string]*string) *ListAncestorsResponse {
	s.Headers = v
	return s
}

func (s *ListAncestorsResponse) SetStatusCode(v int32) *ListAncestorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAncestorsResponse) SetBody(v *ListAncestorsResponseBody) *ListAncestorsResponse {
	s.Body = v
	return s
}

func (s *ListAncestorsResponse) Validate() error {
	return dara.Validate(s)
}
