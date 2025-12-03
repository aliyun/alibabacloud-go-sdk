// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedBranchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProtectedBranchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProtectedBranchesResponse
	GetStatusCode() *int32
	SetBody(v *ListProtectedBranchesResponseBody) *ListProtectedBranchesResponse
	GetBody() *ListProtectedBranchesResponseBody
}

type ListProtectedBranchesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProtectedBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProtectedBranchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProtectedBranchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProtectedBranchesResponse) GetBody() *ListProtectedBranchesResponseBody {
	return s.Body
}

func (s *ListProtectedBranchesResponse) SetHeaders(v map[string]*string) *ListProtectedBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListProtectedBranchesResponse) SetStatusCode(v int32) *ListProtectedBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProtectedBranchesResponse) SetBody(v *ListProtectedBranchesResponseBody) *ListProtectedBranchesResponse {
	s.Body = v
	return s
}

func (s *ListProtectedBranchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
