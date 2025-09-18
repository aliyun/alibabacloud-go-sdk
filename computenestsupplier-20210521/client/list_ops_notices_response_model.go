// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsNoticesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpsNoticesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpsNoticesResponse
	GetStatusCode() *int32
	SetBody(v *ListOpsNoticesResponseBody) *ListOpsNoticesResponse
	GetBody() *ListOpsNoticesResponseBody
}

type ListOpsNoticesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpsNoticesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpsNoticesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpsNoticesResponse) GoString() string {
	return s.String()
}

func (s *ListOpsNoticesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpsNoticesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpsNoticesResponse) GetBody() *ListOpsNoticesResponseBody {
	return s.Body
}

func (s *ListOpsNoticesResponse) SetHeaders(v map[string]*string) *ListOpsNoticesResponse {
	s.Headers = v
	return s
}

func (s *ListOpsNoticesResponse) SetStatusCode(v int32) *ListOpsNoticesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpsNoticesResponse) SetBody(v *ListOpsNoticesResponseBody) *ListOpsNoticesResponse {
	s.Body = v
	return s
}

func (s *ListOpsNoticesResponse) Validate() error {
	return dara.Validate(s)
}
