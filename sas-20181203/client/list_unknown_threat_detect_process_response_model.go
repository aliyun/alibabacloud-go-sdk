// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnknownThreatDetectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnknownThreatDetectProcessResponse
	GetStatusCode() *int32
	SetBody(v *ListUnknownThreatDetectProcessResponseBody) *ListUnknownThreatDetectProcessResponse
	GetBody() *ListUnknownThreatDetectProcessResponseBody
}

type ListUnknownThreatDetectProcessResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnknownThreatDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnknownThreatDetectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnknownThreatDetectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnknownThreatDetectProcessResponse) GetBody() *ListUnknownThreatDetectProcessResponseBody {
	return s.Body
}

func (s *ListUnknownThreatDetectProcessResponse) SetHeaders(v map[string]*string) *ListUnknownThreatDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *ListUnknownThreatDetectProcessResponse) SetStatusCode(v int32) *ListUnknownThreatDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponse) SetBody(v *ListUnknownThreatDetectProcessResponseBody) *ListUnknownThreatDetectProcessResponse {
	s.Body = v
	return s
}

func (s *ListUnknownThreatDetectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
