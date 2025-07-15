// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmsMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmsMetadataResponse
	GetStatusCode() *int32
	SetBody(v *ListSmsMetadataResponseBody) *ListSmsMetadataResponse
	GetBody() *ListSmsMetadataResponseBody
}

type ListSmsMetadataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmsMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmsMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmsMetadataResponse) GoString() string {
	return s.String()
}

func (s *ListSmsMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmsMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmsMetadataResponse) GetBody() *ListSmsMetadataResponseBody {
	return s.Body
}

func (s *ListSmsMetadataResponse) SetHeaders(v map[string]*string) *ListSmsMetadataResponse {
	s.Headers = v
	return s
}

func (s *ListSmsMetadataResponse) SetStatusCode(v int32) *ListSmsMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmsMetadataResponse) SetBody(v *ListSmsMetadataResponseBody) *ListSmsMetadataResponse {
	s.Body = v
	return s
}

func (s *ListSmsMetadataResponse) Validate() error {
	return dara.Validate(s)
}
