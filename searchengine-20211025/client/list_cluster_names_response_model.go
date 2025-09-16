// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterNamesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterNamesResponseBody) *ListClusterNamesResponse
	GetBody() *ListClusterNamesResponseBody
}

type ListClusterNamesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNamesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterNamesResponse) GetBody() *ListClusterNamesResponseBody {
	return s.Body
}

func (s *ListClusterNamesResponse) SetHeaders(v map[string]*string) *ListClusterNamesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNamesResponse) SetStatusCode(v int32) *ListClusterNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNamesResponse) SetBody(v *ListClusterNamesResponseBody) *ListClusterNamesResponse {
	s.Body = v
	return s
}

func (s *ListClusterNamesResponse) Validate() error {
	return dara.Validate(s)
}
