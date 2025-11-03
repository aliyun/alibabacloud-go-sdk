// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorFilterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficMirrorFilterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficMirrorFilterRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficMirrorFilterRulesResponseBody) *CreateTrafficMirrorFilterRulesResponse
	GetBody() *CreateTrafficMirrorFilterRulesResponseBody
}

type CreateTrafficMirrorFilterRulesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficMirrorFilterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficMirrorFilterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficMirrorFilterRulesResponse) GetBody() *CreateTrafficMirrorFilterRulesResponseBody {
	return s.Body
}

func (s *CreateTrafficMirrorFilterRulesResponse) SetHeaders(v map[string]*string) *CreateTrafficMirrorFilterRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponse) SetStatusCode(v int32) *CreateTrafficMirrorFilterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponse) SetBody(v *CreateTrafficMirrorFilterRulesResponseBody) *CreateTrafficMirrorFilterRulesResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
