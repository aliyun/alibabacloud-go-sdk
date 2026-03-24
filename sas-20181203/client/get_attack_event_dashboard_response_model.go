// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackEventDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackEventDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackEventDashboardResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackEventDashboardResponseBody) *GetAttackEventDashboardResponse
	GetBody() *GetAttackEventDashboardResponseBody
}

type GetAttackEventDashboardResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackEventDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackEventDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackEventDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetAttackEventDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackEventDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackEventDashboardResponse) GetBody() *GetAttackEventDashboardResponseBody {
	return s.Body
}

func (s *GetAttackEventDashboardResponse) SetHeaders(v map[string]*string) *GetAttackEventDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetAttackEventDashboardResponse) SetStatusCode(v int32) *GetAttackEventDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackEventDashboardResponse) SetBody(v *GetAttackEventDashboardResponseBody) *GetAttackEventDashboardResponse {
	s.Body = v
	return s
}

func (s *GetAttackEventDashboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
