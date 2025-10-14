// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpResponseHeaderModificationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpResponseHeaderModificationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpResponseHeaderModificationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpResponseHeaderModificationRulesResponseBody) *ListHttpResponseHeaderModificationRulesResponse
	GetBody() *ListHttpResponseHeaderModificationRulesResponseBody
}

type ListHttpResponseHeaderModificationRulesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpResponseHeaderModificationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpResponseHeaderModificationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpResponseHeaderModificationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListHttpResponseHeaderModificationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpResponseHeaderModificationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpResponseHeaderModificationRulesResponse) GetBody() *ListHttpResponseHeaderModificationRulesResponseBody {
	return s.Body
}

func (s *ListHttpResponseHeaderModificationRulesResponse) SetHeaders(v map[string]*string) *ListHttpResponseHeaderModificationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponse) SetStatusCode(v int32) *ListHttpResponseHeaderModificationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponse) SetBody(v *ListHttpResponseHeaderModificationRulesResponseBody) *ListHttpResponseHeaderModificationRulesResponse {
	s.Body = v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
