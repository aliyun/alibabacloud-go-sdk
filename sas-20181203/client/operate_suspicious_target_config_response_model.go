// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSuspiciousTargetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateSuspiciousTargetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateSuspiciousTargetConfigResponse
	GetStatusCode() *int32
	SetBody(v *OperateSuspiciousTargetConfigResponseBody) *OperateSuspiciousTargetConfigResponse
	GetBody() *OperateSuspiciousTargetConfigResponseBody
}

type OperateSuspiciousTargetConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateSuspiciousTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateSuspiciousTargetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateSuspiciousTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousTargetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateSuspiciousTargetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateSuspiciousTargetConfigResponse) GetBody() *OperateSuspiciousTargetConfigResponseBody {
	return s.Body
}

func (s *OperateSuspiciousTargetConfigResponse) SetHeaders(v map[string]*string) *OperateSuspiciousTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateSuspiciousTargetConfigResponse) SetStatusCode(v int32) *OperateSuspiciousTargetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateSuspiciousTargetConfigResponse) SetBody(v *OperateSuspiciousTargetConfigResponseBody) *OperateSuspiciousTargetConfigResponse {
	s.Body = v
	return s
}

func (s *OperateSuspiciousTargetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
