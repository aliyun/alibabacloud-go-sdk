// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnknownThreatDetectStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnknownThreatDetectStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListUnknownThreatDetectStrategyResponseBody) *ListUnknownThreatDetectStrategyResponse
	GetBody() *ListUnknownThreatDetectStrategyResponseBody
}

type ListUnknownThreatDetectStrategyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnknownThreatDetectStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnknownThreatDetectStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnknownThreatDetectStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnknownThreatDetectStrategyResponse) GetBody() *ListUnknownThreatDetectStrategyResponseBody {
	return s.Body
}

func (s *ListUnknownThreatDetectStrategyResponse) SetHeaders(v map[string]*string) *ListUnknownThreatDetectStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponse) SetStatusCode(v int32) *ListUnknownThreatDetectStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponse) SetBody(v *ListUnknownThreatDetectStrategyResponseBody) *ListUnknownThreatDetectStrategyResponse {
	s.Body = v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
