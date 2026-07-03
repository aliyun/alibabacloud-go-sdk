// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResponseRuleFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResponseRuleFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResponseRuleFieldsResponse
	GetStatusCode() *int32
	SetBody(v *ListResponseRuleFieldsResponseBody) *ListResponseRuleFieldsResponse
	GetBody() *ListResponseRuleFieldsResponseBody
}

type ListResponseRuleFieldsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResponseRuleFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResponseRuleFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRuleFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListResponseRuleFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResponseRuleFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResponseRuleFieldsResponse) GetBody() *ListResponseRuleFieldsResponseBody {
	return s.Body
}

func (s *ListResponseRuleFieldsResponse) SetHeaders(v map[string]*string) *ListResponseRuleFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListResponseRuleFieldsResponse) SetStatusCode(v int32) *ListResponseRuleFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResponseRuleFieldsResponse) SetBody(v *ListResponseRuleFieldsResponseBody) *ListResponseRuleFieldsResponse {
	s.Body = v
	return s
}

func (s *ListResponseRuleFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
