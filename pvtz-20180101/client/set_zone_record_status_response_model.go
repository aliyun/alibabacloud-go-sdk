// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetZoneRecordStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetZoneRecordStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetZoneRecordStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetZoneRecordStatusResponseBody) *SetZoneRecordStatusResponse
	GetBody() *SetZoneRecordStatusResponseBody
}

type SetZoneRecordStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetZoneRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetZoneRecordStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetZoneRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetZoneRecordStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetZoneRecordStatusResponse) GetBody() *SetZoneRecordStatusResponseBody {
	return s.Body
}

func (s *SetZoneRecordStatusResponse) SetHeaders(v map[string]*string) *SetZoneRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *SetZoneRecordStatusResponse) SetStatusCode(v int32) *SetZoneRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetZoneRecordStatusResponse) SetBody(v *SetZoneRecordStatusResponseBody) *SetZoneRecordStatusResponse {
	s.Body = v
	return s
}

func (s *SetZoneRecordStatusResponse) Validate() error {
	return dara.Validate(s)
}
