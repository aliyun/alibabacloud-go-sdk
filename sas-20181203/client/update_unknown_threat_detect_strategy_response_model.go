// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnknownThreatDetectStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUnknownThreatDetectStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUnknownThreatDetectStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUnknownThreatDetectStrategyResponseBody) *UpdateUnknownThreatDetectStrategyResponse
	GetBody() *UpdateUnknownThreatDetectStrategyResponseBody
}

type UpdateUnknownThreatDetectStrategyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUnknownThreatDetectStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUnknownThreatDetectStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnknownThreatDetectStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateUnknownThreatDetectStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUnknownThreatDetectStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUnknownThreatDetectStrategyResponse) GetBody() *UpdateUnknownThreatDetectStrategyResponseBody {
	return s.Body
}

func (s *UpdateUnknownThreatDetectStrategyResponse) SetHeaders(v map[string]*string) *UpdateUnknownThreatDetectStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyResponse) SetStatusCode(v int32) *UpdateUnknownThreatDetectStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyResponse) SetBody(v *UpdateUnknownThreatDetectStrategyResponseBody) *UpdateUnknownThreatDetectStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateUnknownThreatDetectStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
