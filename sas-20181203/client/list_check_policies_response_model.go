// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckPoliciesResponseBody) *ListCheckPoliciesResponse
	GetBody() *ListCheckPoliciesResponseBody
}

type ListCheckPoliciesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCheckPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckPoliciesResponse) GetBody() *ListCheckPoliciesResponseBody {
	return s.Body
}

func (s *ListCheckPoliciesResponse) SetHeaders(v map[string]*string) *ListCheckPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCheckPoliciesResponse) SetStatusCode(v int32) *ListCheckPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckPoliciesResponse) SetBody(v *ListCheckPoliciesResponseBody) *ListCheckPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListCheckPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
