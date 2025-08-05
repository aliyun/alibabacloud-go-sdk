// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetJobResourceUsageResponseBody) *GetJobResourceUsageResponse
	GetBody() *GetJobResourceUsageResponseBody
}

type GetJobResourceUsageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobResourceUsageResponse) GetBody() *GetJobResourceUsageResponseBody {
	return s.Body
}

func (s *GetJobResourceUsageResponse) SetHeaders(v map[string]*string) *GetJobResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetJobResourceUsageResponse) SetStatusCode(v int32) *GetJobResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResourceUsageResponse) SetBody(v *GetJobResourceUsageResponseBody) *GetJobResourceUsageResponse {
	s.Body = v
	return s
}

func (s *GetJobResourceUsageResponse) Validate() error {
	return dara.Validate(s)
}
