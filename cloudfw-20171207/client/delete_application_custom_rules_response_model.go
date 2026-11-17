// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationCustomRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationCustomRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationCustomRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationCustomRulesResponseBody) *DeleteApplicationCustomRulesResponse
	GetBody() *DeleteApplicationCustomRulesResponseBody
}

type DeleteApplicationCustomRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationCustomRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationCustomRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationCustomRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationCustomRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationCustomRulesResponse) GetBody() *DeleteApplicationCustomRulesResponseBody {
	return s.Body
}

func (s *DeleteApplicationCustomRulesResponse) SetHeaders(v map[string]*string) *DeleteApplicationCustomRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationCustomRulesResponse) SetStatusCode(v int32) *DeleteApplicationCustomRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationCustomRulesResponse) SetBody(v *DeleteApplicationCustomRulesResponseBody) *DeleteApplicationCustomRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationCustomRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
