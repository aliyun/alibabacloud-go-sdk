// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnknownThreatDetectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUnknownThreatDetectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUnknownThreatDetectProcessResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUnknownThreatDetectProcessResponseBody) *UpdateUnknownThreatDetectProcessResponse
	GetBody() *UpdateUnknownThreatDetectProcessResponseBody
}

type UpdateUnknownThreatDetectProcessResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUnknownThreatDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUnknownThreatDetectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnknownThreatDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *UpdateUnknownThreatDetectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUnknownThreatDetectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUnknownThreatDetectProcessResponse) GetBody() *UpdateUnknownThreatDetectProcessResponseBody {
	return s.Body
}

func (s *UpdateUnknownThreatDetectProcessResponse) SetHeaders(v map[string]*string) *UpdateUnknownThreatDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *UpdateUnknownThreatDetectProcessResponse) SetStatusCode(v int32) *UpdateUnknownThreatDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUnknownThreatDetectProcessResponse) SetBody(v *UpdateUnknownThreatDetectProcessResponseBody) *UpdateUnknownThreatDetectProcessResponse {
	s.Body = v
	return s
}

func (s *UpdateUnknownThreatDetectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
