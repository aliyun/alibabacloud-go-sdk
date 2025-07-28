// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckZoneNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckZoneNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckZoneNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckZoneNameResponseBody) *CheckZoneNameResponse
	GetBody() *CheckZoneNameResponseBody
}

type CheckZoneNameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckZoneNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckZoneNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckZoneNameResponse) GoString() string {
	return s.String()
}

func (s *CheckZoneNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckZoneNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckZoneNameResponse) GetBody() *CheckZoneNameResponseBody {
	return s.Body
}

func (s *CheckZoneNameResponse) SetHeaders(v map[string]*string) *CheckZoneNameResponse {
	s.Headers = v
	return s
}

func (s *CheckZoneNameResponse) SetStatusCode(v int32) *CheckZoneNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckZoneNameResponse) SetBody(v *CheckZoneNameResponseBody) *CheckZoneNameResponse {
	s.Body = v
	return s
}

func (s *CheckZoneNameResponse) Validate() error {
	return dara.Validate(s)
}
