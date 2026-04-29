// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleUnknownThreatDetectEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleUnknownThreatDetectEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleUnknownThreatDetectEventResponse
	GetStatusCode() *int32
	SetBody(v *HandleUnknownThreatDetectEventResponseBody) *HandleUnknownThreatDetectEventResponse
	GetBody() *HandleUnknownThreatDetectEventResponseBody
}

type HandleUnknownThreatDetectEventResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleUnknownThreatDetectEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleUnknownThreatDetectEventResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleUnknownThreatDetectEventResponse) GoString() string {
	return s.String()
}

func (s *HandleUnknownThreatDetectEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleUnknownThreatDetectEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleUnknownThreatDetectEventResponse) GetBody() *HandleUnknownThreatDetectEventResponseBody {
	return s.Body
}

func (s *HandleUnknownThreatDetectEventResponse) SetHeaders(v map[string]*string) *HandleUnknownThreatDetectEventResponse {
	s.Headers = v
	return s
}

func (s *HandleUnknownThreatDetectEventResponse) SetStatusCode(v int32) *HandleUnknownThreatDetectEventResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleUnknownThreatDetectEventResponse) SetBody(v *HandleUnknownThreatDetectEventResponseBody) *HandleUnknownThreatDetectEventResponse {
	s.Body = v
	return s
}

func (s *HandleUnknownThreatDetectEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
