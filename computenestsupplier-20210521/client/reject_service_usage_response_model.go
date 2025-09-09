// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectServiceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectServiceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectServiceUsageResponse
	GetStatusCode() *int32
	SetBody(v *RejectServiceUsageResponseBody) *RejectServiceUsageResponse
	GetBody() *RejectServiceUsageResponseBody
}

type RejectServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectServiceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *RejectServiceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectServiceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectServiceUsageResponse) GetBody() *RejectServiceUsageResponseBody {
	return s.Body
}

func (s *RejectServiceUsageResponse) SetHeaders(v map[string]*string) *RejectServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *RejectServiceUsageResponse) SetStatusCode(v int32) *RejectServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectServiceUsageResponse) SetBody(v *RejectServiceUsageResponseBody) *RejectServiceUsageResponse {
	s.Body = v
	return s
}

func (s *RejectServiceUsageResponse) Validate() error {
	return dara.Validate(s)
}
