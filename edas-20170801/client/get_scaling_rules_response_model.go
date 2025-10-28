// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScalingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScalingRulesResponse
	GetStatusCode() *int32
	SetBody(v *GetScalingRulesResponseBody) *GetScalingRulesResponse
	GetBody() *GetScalingRulesResponseBody
}

type GetScalingRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScalingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScalingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScalingRulesResponse) GetBody() *GetScalingRulesResponseBody {
	return s.Body
}

func (s *GetScalingRulesResponse) SetHeaders(v map[string]*string) *GetScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *GetScalingRulesResponse) SetStatusCode(v int32) *GetScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScalingRulesResponse) SetBody(v *GetScalingRulesResponseBody) *GetScalingRulesResponse {
	s.Body = v
	return s
}

func (s *GetScalingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
