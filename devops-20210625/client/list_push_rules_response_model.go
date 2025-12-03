// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPushRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPushRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPushRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListPushRulesResponseBody) *ListPushRulesResponse
	GetBody() *ListPushRulesResponseBody
}

type ListPushRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPushRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPushRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPushRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPushRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPushRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPushRulesResponse) GetBody() *ListPushRulesResponseBody {
	return s.Body
}

func (s *ListPushRulesResponse) SetHeaders(v map[string]*string) *ListPushRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPushRulesResponse) SetStatusCode(v int32) *ListPushRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPushRulesResponse) SetBody(v *ListPushRulesResponseBody) *ListPushRulesResponse {
	s.Body = v
	return s
}

func (s *ListPushRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
