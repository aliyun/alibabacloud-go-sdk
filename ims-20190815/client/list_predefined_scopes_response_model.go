// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedScopesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPredefinedScopesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPredefinedScopesResponse
	GetStatusCode() *int32
	SetBody(v *ListPredefinedScopesResponseBody) *ListPredefinedScopesResponse
	GetBody() *ListPredefinedScopesResponseBody
}

type ListPredefinedScopesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPredefinedScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPredefinedScopesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedScopesResponse) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPredefinedScopesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPredefinedScopesResponse) GetBody() *ListPredefinedScopesResponseBody {
	return s.Body
}

func (s *ListPredefinedScopesResponse) SetHeaders(v map[string]*string) *ListPredefinedScopesResponse {
	s.Headers = v
	return s
}

func (s *ListPredefinedScopesResponse) SetStatusCode(v int32) *ListPredefinedScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPredefinedScopesResponse) SetBody(v *ListPredefinedScopesResponseBody) *ListPredefinedScopesResponse {
	s.Body = v
	return s
}

func (s *ListPredefinedScopesResponse) Validate() error {
	return dara.Validate(s)
}
