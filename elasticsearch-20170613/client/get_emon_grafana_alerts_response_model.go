// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonGrafanaAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmonGrafanaAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmonGrafanaAlertsResponse
	GetStatusCode() *int32
	SetBody(v *GetEmonGrafanaAlertsResponseBody) *GetEmonGrafanaAlertsResponse
	GetBody() *GetEmonGrafanaAlertsResponseBody
}

type GetEmonGrafanaAlertsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmonGrafanaAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonGrafanaAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmonGrafanaAlertsResponse) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmonGrafanaAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmonGrafanaAlertsResponse) GetBody() *GetEmonGrafanaAlertsResponseBody {
	return s.Body
}

func (s *GetEmonGrafanaAlertsResponse) SetHeaders(v map[string]*string) *GetEmonGrafanaAlertsResponse {
	s.Headers = v
	return s
}

func (s *GetEmonGrafanaAlertsResponse) SetStatusCode(v int32) *GetEmonGrafanaAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponse) SetBody(v *GetEmonGrafanaAlertsResponseBody) *GetEmonGrafanaAlertsResponse {
	s.Body = v
	return s
}

func (s *GetEmonGrafanaAlertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
