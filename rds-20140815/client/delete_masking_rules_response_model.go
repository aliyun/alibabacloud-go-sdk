// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaskingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaskingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaskingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaskingRulesResponseBody) *DeleteMaskingRulesResponse
	GetBody() *DeleteMaskingRulesResponseBody
}

type DeleteMaskingRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaskingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaskingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaskingRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaskingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaskingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaskingRulesResponse) GetBody() *DeleteMaskingRulesResponseBody {
	return s.Body
}

func (s *DeleteMaskingRulesResponse) SetHeaders(v map[string]*string) *DeleteMaskingRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaskingRulesResponse) SetStatusCode(v int32) *DeleteMaskingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaskingRulesResponse) SetBody(v *DeleteMaskingRulesResponseBody) *DeleteMaskingRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteMaskingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
