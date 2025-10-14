// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectionRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDetectionRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDetectionRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListDetectionRulesResponseBody) *ListDetectionRulesResponse
	GetBody() *ListDetectionRulesResponseBody
}

type ListDetectionRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDetectionRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDetectionRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDetectionRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDetectionRulesResponse) GetBody() *ListDetectionRulesResponseBody {
	return s.Body
}

func (s *ListDetectionRulesResponse) SetHeaders(v map[string]*string) *ListDetectionRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDetectionRulesResponse) SetStatusCode(v int32) *ListDetectionRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDetectionRulesResponse) SetBody(v *ListDetectionRulesResponseBody) *ListDetectionRulesResponse {
	s.Body = v
	return s
}

func (s *ListDetectionRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
