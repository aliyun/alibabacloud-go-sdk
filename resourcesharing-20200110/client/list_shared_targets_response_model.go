// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSharedTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSharedTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ListSharedTargetsResponseBody) *ListSharedTargetsResponse
	GetBody() *ListSharedTargetsResponseBody
}

type ListSharedTargetsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSharedTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSharedTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSharedTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSharedTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSharedTargetsResponse) GetBody() *ListSharedTargetsResponseBody {
	return s.Body
}

func (s *ListSharedTargetsResponse) SetHeaders(v map[string]*string) *ListSharedTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListSharedTargetsResponse) SetStatusCode(v int32) *ListSharedTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedTargetsResponse) SetBody(v *ListSharedTargetsResponseBody) *ListSharedTargetsResponse {
	s.Body = v
	return s
}

func (s *ListSharedTargetsResponse) Validate() error {
	return dara.Validate(s)
}
