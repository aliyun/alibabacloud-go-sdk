// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUnknownThreatDetectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUnknownThreatDetectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUnknownThreatDetectProcessResponse
	GetStatusCode() *int32
	SetBody(v *AddUnknownThreatDetectProcessResponseBody) *AddUnknownThreatDetectProcessResponse
	GetBody() *AddUnknownThreatDetectProcessResponseBody
}

type AddUnknownThreatDetectProcessResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUnknownThreatDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUnknownThreatDetectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUnknownThreatDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *AddUnknownThreatDetectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUnknownThreatDetectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUnknownThreatDetectProcessResponse) GetBody() *AddUnknownThreatDetectProcessResponseBody {
	return s.Body
}

func (s *AddUnknownThreatDetectProcessResponse) SetHeaders(v map[string]*string) *AddUnknownThreatDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *AddUnknownThreatDetectProcessResponse) SetStatusCode(v int32) *AddUnknownThreatDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUnknownThreatDetectProcessResponse) SetBody(v *AddUnknownThreatDetectProcessResponseBody) *AddUnknownThreatDetectProcessResponse {
	s.Body = v
	return s
}

func (s *AddUnknownThreatDetectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
