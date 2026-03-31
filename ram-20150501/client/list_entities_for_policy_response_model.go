// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesForPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEntitiesForPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEntitiesForPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListEntitiesForPolicyResponseBody) *ListEntitiesForPolicyResponse
	GetBody() *ListEntitiesForPolicyResponseBody
}

type ListEntitiesForPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesForPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesForPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEntitiesForPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEntitiesForPolicyResponse) GetBody() *ListEntitiesForPolicyResponseBody {
	return s.Body
}

func (s *ListEntitiesForPolicyResponse) SetHeaders(v map[string]*string) *ListEntitiesForPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesForPolicyResponse) SetStatusCode(v int32) *ListEntitiesForPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesForPolicyResponse) SetBody(v *ListEntitiesForPolicyResponseBody) *ListEntitiesForPolicyResponse {
	s.Body = v
	return s
}

func (s *ListEntitiesForPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
