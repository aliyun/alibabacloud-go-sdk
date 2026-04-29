// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnknownThreatDetectEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnknownThreatDetectEventResponse
	GetStatusCode() *int32
	SetBody(v *ListUnknownThreatDetectEventResponseBody) *ListUnknownThreatDetectEventResponse
	GetBody() *ListUnknownThreatDetectEventResponseBody
}

type ListUnknownThreatDetectEventResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnknownThreatDetectEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnknownThreatDetectEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectEventResponse) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnknownThreatDetectEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnknownThreatDetectEventResponse) GetBody() *ListUnknownThreatDetectEventResponseBody {
	return s.Body
}

func (s *ListUnknownThreatDetectEventResponse) SetHeaders(v map[string]*string) *ListUnknownThreatDetectEventResponse {
	s.Headers = v
	return s
}

func (s *ListUnknownThreatDetectEventResponse) SetStatusCode(v int32) *ListUnknownThreatDetectEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponse) SetBody(v *ListUnknownThreatDetectEventResponseBody) *ListUnknownThreatDetectEventResponse {
	s.Body = v
	return s
}

func (s *ListUnknownThreatDetectEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
