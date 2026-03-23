// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepairApRadioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RepairApRadioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RepairApRadioResponse
	GetStatusCode() *int32
	SetBody(v *RepairApRadioResponseBody) *RepairApRadioResponse
	GetBody() *RepairApRadioResponseBody
}

type RepairApRadioResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RepairApRadioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RepairApRadioResponse) String() string {
	return dara.Prettify(s)
}

func (s RepairApRadioResponse) GoString() string {
	return s.String()
}

func (s *RepairApRadioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RepairApRadioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RepairApRadioResponse) GetBody() *RepairApRadioResponseBody {
	return s.Body
}

func (s *RepairApRadioResponse) SetHeaders(v map[string]*string) *RepairApRadioResponse {
	s.Headers = v
	return s
}

func (s *RepairApRadioResponse) SetStatusCode(v int32) *RepairApRadioResponse {
	s.StatusCode = &v
	return s
}

func (s *RepairApRadioResponse) SetBody(v *RepairApRadioResponseBody) *RepairApRadioResponse {
	s.Body = v
	return s
}

func (s *RepairApRadioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
