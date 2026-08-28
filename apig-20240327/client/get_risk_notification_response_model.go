// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRiskNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRiskNotificationResponse
	GetStatusCode() *int32
	SetBody(v *GetRiskNotificationResponseBody) *GetRiskNotificationResponse
	GetBody() *GetRiskNotificationResponseBody
}

type GetRiskNotificationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRiskNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRiskNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRiskNotificationResponse) GoString() string {
	return s.String()
}

func (s *GetRiskNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRiskNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRiskNotificationResponse) GetBody() *GetRiskNotificationResponseBody {
	return s.Body
}

func (s *GetRiskNotificationResponse) SetHeaders(v map[string]*string) *GetRiskNotificationResponse {
	s.Headers = v
	return s
}

func (s *GetRiskNotificationResponse) SetStatusCode(v int32) *GetRiskNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRiskNotificationResponse) SetBody(v *GetRiskNotificationResponseBody) *GetRiskNotificationResponse {
	s.Body = v
	return s
}

func (s *GetRiskNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
