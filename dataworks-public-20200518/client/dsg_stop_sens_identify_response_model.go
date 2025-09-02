// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgStopSensIdentifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgStopSensIdentifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgStopSensIdentifyResponse
	GetStatusCode() *int32
	SetBody(v *DsgStopSensIdentifyResponseBody) *DsgStopSensIdentifyResponse
	GetBody() *DsgStopSensIdentifyResponseBody
}

type DsgStopSensIdentifyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgStopSensIdentifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgStopSensIdentifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgStopSensIdentifyResponse) GoString() string {
	return s.String()
}

func (s *DsgStopSensIdentifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgStopSensIdentifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgStopSensIdentifyResponse) GetBody() *DsgStopSensIdentifyResponseBody {
	return s.Body
}

func (s *DsgStopSensIdentifyResponse) SetHeaders(v map[string]*string) *DsgStopSensIdentifyResponse {
	s.Headers = v
	return s
}

func (s *DsgStopSensIdentifyResponse) SetStatusCode(v int32) *DsgStopSensIdentifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgStopSensIdentifyResponse) SetBody(v *DsgStopSensIdentifyResponseBody) *DsgStopSensIdentifyResponse {
	s.Body = v
	return s
}

func (s *DsgStopSensIdentifyResponse) Validate() error {
	return dara.Validate(s)
}
