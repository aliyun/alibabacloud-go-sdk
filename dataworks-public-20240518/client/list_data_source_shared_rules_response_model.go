// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceSharedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceSharedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceSharedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceSharedRulesResponseBody) *ListDataSourceSharedRulesResponse
	GetBody() *ListDataSourceSharedRulesResponseBody
}

type ListDataSourceSharedRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceSharedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceSharedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSharedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceSharedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceSharedRulesResponse) GetBody() *ListDataSourceSharedRulesResponseBody {
	return s.Body
}

func (s *ListDataSourceSharedRulesResponse) SetHeaders(v map[string]*string) *ListDataSourceSharedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceSharedRulesResponse) SetStatusCode(v int32) *ListDataSourceSharedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceSharedRulesResponse) SetBody(v *ListDataSourceSharedRulesResponseBody) *ListDataSourceSharedRulesResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceSharedRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
