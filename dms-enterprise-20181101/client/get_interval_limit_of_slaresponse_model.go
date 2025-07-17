// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntervalLimitOfSLAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntervalLimitOfSLAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntervalLimitOfSLAResponse
	GetStatusCode() *int32
	SetBody(v *GetIntervalLimitOfSLAResponseBody) *GetIntervalLimitOfSLAResponse
	GetBody() *GetIntervalLimitOfSLAResponseBody
}

type GetIntervalLimitOfSLAResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntervalLimitOfSLAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntervalLimitOfSLAResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntervalLimitOfSLAResponse) GoString() string {
	return s.String()
}

func (s *GetIntervalLimitOfSLAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntervalLimitOfSLAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntervalLimitOfSLAResponse) GetBody() *GetIntervalLimitOfSLAResponseBody {
	return s.Body
}

func (s *GetIntervalLimitOfSLAResponse) SetHeaders(v map[string]*string) *GetIntervalLimitOfSLAResponse {
	s.Headers = v
	return s
}

func (s *GetIntervalLimitOfSLAResponse) SetStatusCode(v int32) *GetIntervalLimitOfSLAResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntervalLimitOfSLAResponse) SetBody(v *GetIntervalLimitOfSLAResponseBody) *GetIntervalLimitOfSLAResponse {
	s.Body = v
	return s
}

func (s *GetIntervalLimitOfSLAResponse) Validate() error {
	return dara.Validate(s)
}
