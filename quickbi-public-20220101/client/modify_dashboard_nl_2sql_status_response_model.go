// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDashboardNl2sqlStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDashboardNl2sqlStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDashboardNl2sqlStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDashboardNl2sqlStatusResponseBody) *ModifyDashboardNl2sqlStatusResponse
	GetBody() *ModifyDashboardNl2sqlStatusResponseBody
}

type ModifyDashboardNl2sqlStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDashboardNl2sqlStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDashboardNl2sqlStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDashboardNl2sqlStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDashboardNl2sqlStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDashboardNl2sqlStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDashboardNl2sqlStatusResponse) GetBody() *ModifyDashboardNl2sqlStatusResponseBody {
	return s.Body
}

func (s *ModifyDashboardNl2sqlStatusResponse) SetHeaders(v map[string]*string) *ModifyDashboardNl2sqlStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDashboardNl2sqlStatusResponse) SetStatusCode(v int32) *ModifyDashboardNl2sqlStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDashboardNl2sqlStatusResponse) SetBody(v *ModifyDashboardNl2sqlStatusResponseBody) *ModifyDashboardNl2sqlStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyDashboardNl2sqlStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
