// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceUsageResponseBody) *GetInstanceUsageResponse
	GetBody() *GetInstanceUsageResponseBody
}

type GetInstanceUsageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceUsageResponse) GetBody() *GetInstanceUsageResponseBody {
	return s.Body
}

func (s *GetInstanceUsageResponse) SetHeaders(v map[string]*string) *GetInstanceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceUsageResponse) SetStatusCode(v int32) *GetInstanceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceUsageResponse) SetBody(v *GetInstanceUsageResponseBody) *GetInstanceUsageResponse {
	s.Body = v
	return s
}

func (s *GetInstanceUsageResponse) Validate() error {
	return dara.Validate(s)
}
