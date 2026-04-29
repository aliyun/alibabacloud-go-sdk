// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMOQuotaAlertThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMOQuotaAlertThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMOQuotaAlertThresholdResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMOQuotaAlertThresholdResponseBody) *UpdateMOQuotaAlertThresholdResponse
	GetBody() *UpdateMOQuotaAlertThresholdResponseBody
}

type UpdateMOQuotaAlertThresholdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMOQuotaAlertThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMOQuotaAlertThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMOQuotaAlertThresholdResponse) GoString() string {
	return s.String()
}

func (s *UpdateMOQuotaAlertThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMOQuotaAlertThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMOQuotaAlertThresholdResponse) GetBody() *UpdateMOQuotaAlertThresholdResponseBody {
	return s.Body
}

func (s *UpdateMOQuotaAlertThresholdResponse) SetHeaders(v map[string]*string) *UpdateMOQuotaAlertThresholdResponse {
	s.Headers = v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponse) SetStatusCode(v int32) *UpdateMOQuotaAlertThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponse) SetBody(v *UpdateMOQuotaAlertThresholdResponseBody) *UpdateMOQuotaAlertThresholdResponse {
	s.Body = v
	return s
}

func (s *UpdateMOQuotaAlertThresholdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
