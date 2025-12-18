// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *CopyConfigRulesResponseBody) *CopyConfigRulesResponse
	GetBody() *CopyConfigRulesResponseBody
}

type CopyConfigRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *CopyConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyConfigRulesResponse) GetBody() *CopyConfigRulesResponseBody {
	return s.Body
}

func (s *CopyConfigRulesResponse) SetHeaders(v map[string]*string) *CopyConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *CopyConfigRulesResponse) SetStatusCode(v int32) *CopyConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyConfigRulesResponse) SetBody(v *CopyConfigRulesResponseBody) *CopyConfigRulesResponse {
	s.Body = v
	return s
}

func (s *CopyConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
