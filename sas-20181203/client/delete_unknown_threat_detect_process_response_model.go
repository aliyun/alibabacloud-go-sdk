// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnknownThreatDetectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUnknownThreatDetectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUnknownThreatDetectProcessResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUnknownThreatDetectProcessResponseBody) *DeleteUnknownThreatDetectProcessResponse
	GetBody() *DeleteUnknownThreatDetectProcessResponseBody
}

type DeleteUnknownThreatDetectProcessResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUnknownThreatDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUnknownThreatDetectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnknownThreatDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteUnknownThreatDetectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUnknownThreatDetectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUnknownThreatDetectProcessResponse) GetBody() *DeleteUnknownThreatDetectProcessResponseBody {
	return s.Body
}

func (s *DeleteUnknownThreatDetectProcessResponse) SetHeaders(v map[string]*string) *DeleteUnknownThreatDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteUnknownThreatDetectProcessResponse) SetStatusCode(v int32) *DeleteUnknownThreatDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUnknownThreatDetectProcessResponse) SetBody(v *DeleteUnknownThreatDetectProcessResponseBody) *DeleteUnknownThreatDetectProcessResponse {
	s.Body = v
	return s
}

func (s *DeleteUnknownThreatDetectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
