// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnknownThreatDetectStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUnknownThreatDetectStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUnknownThreatDetectStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateUnknownThreatDetectStrategyResponseBody) *CreateUnknownThreatDetectStrategyResponse
	GetBody() *CreateUnknownThreatDetectStrategyResponseBody
}

type CreateUnknownThreatDetectStrategyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUnknownThreatDetectStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUnknownThreatDetectStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUnknownThreatDetectStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateUnknownThreatDetectStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUnknownThreatDetectStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUnknownThreatDetectStrategyResponse) GetBody() *CreateUnknownThreatDetectStrategyResponseBody {
	return s.Body
}

func (s *CreateUnknownThreatDetectStrategyResponse) SetHeaders(v map[string]*string) *CreateUnknownThreatDetectStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateUnknownThreatDetectStrategyResponse) SetStatusCode(v int32) *CreateUnknownThreatDetectStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUnknownThreatDetectStrategyResponse) SetBody(v *CreateUnknownThreatDetectStrategyResponseBody) *CreateUnknownThreatDetectStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateUnknownThreatDetectStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
