// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRiskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRiskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRiskStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRiskStatusResponseBody) *UpdateRiskStatusResponse
	GetBody() *UpdateRiskStatusResponseBody
}

type UpdateRiskStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRiskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRiskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRiskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateRiskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRiskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRiskStatusResponse) GetBody() *UpdateRiskStatusResponseBody {
	return s.Body
}

func (s *UpdateRiskStatusResponse) SetHeaders(v map[string]*string) *UpdateRiskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateRiskStatusResponse) SetStatusCode(v int32) *UpdateRiskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRiskStatusResponse) SetBody(v *UpdateRiskStatusResponseBody) *UpdateRiskStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateRiskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
