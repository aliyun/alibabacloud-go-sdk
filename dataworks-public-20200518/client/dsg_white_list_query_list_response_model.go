// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListQueryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgWhiteListQueryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgWhiteListQueryListResponse
	GetStatusCode() *int32
	SetBody(v *DsgWhiteListQueryListResponseBody) *DsgWhiteListQueryListResponse
	GetBody() *DsgWhiteListQueryListResponseBody
}

type DsgWhiteListQueryListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgWhiteListQueryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgWhiteListQueryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListQueryListResponse) GoString() string {
	return s.String()
}

func (s *DsgWhiteListQueryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgWhiteListQueryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgWhiteListQueryListResponse) GetBody() *DsgWhiteListQueryListResponseBody {
	return s.Body
}

func (s *DsgWhiteListQueryListResponse) SetHeaders(v map[string]*string) *DsgWhiteListQueryListResponse {
	s.Headers = v
	return s
}

func (s *DsgWhiteListQueryListResponse) SetStatusCode(v int32) *DsgWhiteListQueryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgWhiteListQueryListResponse) SetBody(v *DsgWhiteListQueryListResponseBody) *DsgWhiteListQueryListResponse {
	s.Body = v
	return s
}

func (s *DsgWhiteListQueryListResponse) Validate() error {
	return dara.Validate(s)
}
