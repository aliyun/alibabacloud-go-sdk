// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLiveShiftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenLiveShiftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenLiveShiftResponse
	GetStatusCode() *int32
	SetBody(v *OpenLiveShiftResponseBody) *OpenLiveShiftResponse
	GetBody() *OpenLiveShiftResponseBody
}

type OpenLiveShiftResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenLiveShiftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenLiveShiftResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenLiveShiftResponse) GoString() string {
	return s.String()
}

func (s *OpenLiveShiftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenLiveShiftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenLiveShiftResponse) GetBody() *OpenLiveShiftResponseBody {
	return s.Body
}

func (s *OpenLiveShiftResponse) SetHeaders(v map[string]*string) *OpenLiveShiftResponse {
	s.Headers = v
	return s
}

func (s *OpenLiveShiftResponse) SetStatusCode(v int32) *OpenLiveShiftResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenLiveShiftResponse) SetBody(v *OpenLiveShiftResponseBody) *OpenLiveShiftResponse {
	s.Body = v
	return s
}

func (s *OpenLiveShiftResponse) Validate() error {
	return dara.Validate(s)
}
