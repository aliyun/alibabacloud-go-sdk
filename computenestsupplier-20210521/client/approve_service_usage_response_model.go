// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveServiceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveServiceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveServiceUsageResponse
	GetStatusCode() *int32
	SetBody(v *ApproveServiceUsageResponseBody) *ApproveServiceUsageResponse
	GetBody() *ApproveServiceUsageResponseBody
}

type ApproveServiceUsageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveServiceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *ApproveServiceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveServiceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveServiceUsageResponse) GetBody() *ApproveServiceUsageResponseBody {
	return s.Body
}

func (s *ApproveServiceUsageResponse) SetHeaders(v map[string]*string) *ApproveServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *ApproveServiceUsageResponse) SetStatusCode(v int32) *ApproveServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveServiceUsageResponse) SetBody(v *ApproveServiceUsageResponseBody) *ApproveServiceUsageResponse {
	s.Body = v
	return s
}

func (s *ApproveServiceUsageResponse) Validate() error {
	return dara.Validate(s)
}
