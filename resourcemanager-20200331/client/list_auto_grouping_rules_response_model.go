// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoGroupingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoGroupingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoGroupingRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoGroupingRulesResponseBody) *ListAutoGroupingRulesResponse
	GetBody() *ListAutoGroupingRulesResponseBody
}

type ListAutoGroupingRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoGroupingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoGroupingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoGroupingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoGroupingRulesResponse) GetBody() *ListAutoGroupingRulesResponseBody {
	return s.Body
}

func (s *ListAutoGroupingRulesResponse) SetHeaders(v map[string]*string) *ListAutoGroupingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoGroupingRulesResponse) SetStatusCode(v int32) *ListAutoGroupingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoGroupingRulesResponse) SetBody(v *ListAutoGroupingRulesResponseBody) *ListAutoGroupingRulesResponse {
	s.Body = v
	return s
}

func (s *ListAutoGroupingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
