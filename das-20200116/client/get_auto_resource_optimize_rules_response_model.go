// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoResourceOptimizeRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoResourceOptimizeRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoResourceOptimizeRulesResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoResourceOptimizeRulesResponseBody) *GetAutoResourceOptimizeRulesResponse
	GetBody() *GetAutoResourceOptimizeRulesResponseBody
}

type GetAutoResourceOptimizeRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoResourceOptimizeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoResourceOptimizeRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesResponse) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoResourceOptimizeRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoResourceOptimizeRulesResponse) GetBody() *GetAutoResourceOptimizeRulesResponseBody {
	return s.Body
}

func (s *GetAutoResourceOptimizeRulesResponse) SetHeaders(v map[string]*string) *GetAutoResourceOptimizeRulesResponse {
	s.Headers = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponse) SetStatusCode(v int32) *GetAutoResourceOptimizeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponse) SetBody(v *GetAutoResourceOptimizeRulesResponseBody) *GetAutoResourceOptimizeRulesResponse {
	s.Body = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
