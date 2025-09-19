// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionTargetDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterceptionTargetDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterceptionTargetDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetInterceptionTargetDetailResponseBody) *GetInterceptionTargetDetailResponse
	GetBody() *GetInterceptionTargetDetailResponseBody
}

type GetInterceptionTargetDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterceptionTargetDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterceptionTargetDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionTargetDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInterceptionTargetDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterceptionTargetDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterceptionTargetDetailResponse) GetBody() *GetInterceptionTargetDetailResponseBody {
	return s.Body
}

func (s *GetInterceptionTargetDetailResponse) SetHeaders(v map[string]*string) *GetInterceptionTargetDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInterceptionTargetDetailResponse) SetStatusCode(v int32) *GetInterceptionTargetDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterceptionTargetDetailResponse) SetBody(v *GetInterceptionTargetDetailResponseBody) *GetInterceptionTargetDetailResponse {
	s.Body = v
	return s
}

func (s *GetInterceptionTargetDetailResponse) Validate() error {
	return dara.Validate(s)
}
