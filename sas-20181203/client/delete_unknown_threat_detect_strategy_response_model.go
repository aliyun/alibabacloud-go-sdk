// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnknownThreatDetectStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUnknownThreatDetectStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUnknownThreatDetectStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUnknownThreatDetectStrategyResponseBody) *DeleteUnknownThreatDetectStrategyResponse
	GetBody() *DeleteUnknownThreatDetectStrategyResponseBody
}

type DeleteUnknownThreatDetectStrategyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUnknownThreatDetectStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUnknownThreatDetectStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnknownThreatDetectStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteUnknownThreatDetectStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUnknownThreatDetectStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUnknownThreatDetectStrategyResponse) GetBody() *DeleteUnknownThreatDetectStrategyResponseBody {
	return s.Body
}

func (s *DeleteUnknownThreatDetectStrategyResponse) SetHeaders(v map[string]*string) *DeleteUnknownThreatDetectStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteUnknownThreatDetectStrategyResponse) SetStatusCode(v int32) *DeleteUnknownThreatDetectStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUnknownThreatDetectStrategyResponse) SetBody(v *DeleteUnknownThreatDetectStrategyResponseBody) *DeleteUnknownThreatDetectStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteUnknownThreatDetectStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
