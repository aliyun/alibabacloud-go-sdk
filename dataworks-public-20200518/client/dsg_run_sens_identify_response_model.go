// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgRunSensIdentifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgRunSensIdentifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgRunSensIdentifyResponse
	GetStatusCode() *int32
	SetBody(v *DsgRunSensIdentifyResponseBody) *DsgRunSensIdentifyResponse
	GetBody() *DsgRunSensIdentifyResponseBody
}

type DsgRunSensIdentifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgRunSensIdentifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgRunSensIdentifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgRunSensIdentifyResponse) GoString() string {
	return s.String()
}

func (s *DsgRunSensIdentifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgRunSensIdentifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgRunSensIdentifyResponse) GetBody() *DsgRunSensIdentifyResponseBody {
	return s.Body
}

func (s *DsgRunSensIdentifyResponse) SetHeaders(v map[string]*string) *DsgRunSensIdentifyResponse {
	s.Headers = v
	return s
}

func (s *DsgRunSensIdentifyResponse) SetStatusCode(v int32) *DsgRunSensIdentifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgRunSensIdentifyResponse) SetBody(v *DsgRunSensIdentifyResponseBody) *DsgRunSensIdentifyResponse {
	s.Body = v
	return s
}

func (s *DsgRunSensIdentifyResponse) Validate() error {
	return dara.Validate(s)
}
