// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBranchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBranchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBranchesResponse
	GetStatusCode() *int32
	SetBody(v *ListBranchesResponseBody) *ListBranchesResponse
	GetBody() *ListBranchesResponseBody
}

type ListBranchesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBranchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListBranchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBranchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBranchesResponse) GetBody() *ListBranchesResponseBody {
	return s.Body
}

func (s *ListBranchesResponse) SetHeaders(v map[string]*string) *ListBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListBranchesResponse) SetStatusCode(v int32) *ListBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBranchesResponse) SetBody(v *ListBranchesResponseBody) *ListBranchesResponse {
	s.Body = v
	return s
}

func (s *ListBranchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
