// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorFilterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficMirrorFilterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficMirrorFilterRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficMirrorFilterRulesResponseBody) *DeleteTrafficMirrorFilterRulesResponse
	GetBody() *DeleteTrafficMirrorFilterRulesResponseBody
}

type DeleteTrafficMirrorFilterRulesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficMirrorFilterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficMirrorFilterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorFilterRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorFilterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficMirrorFilterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficMirrorFilterRulesResponse) GetBody() *DeleteTrafficMirrorFilterRulesResponseBody {
	return s.Body
}

func (s *DeleteTrafficMirrorFilterRulesResponse) SetHeaders(v map[string]*string) *DeleteTrafficMirrorFilterRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesResponse) SetStatusCode(v int32) *DeleteTrafficMirrorFilterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesResponse) SetBody(v *DeleteTrafficMirrorFilterRulesResponseBody) *DeleteTrafficMirrorFilterRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
