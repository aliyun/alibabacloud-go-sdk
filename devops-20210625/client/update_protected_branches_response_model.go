// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectedBranchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProtectedBranchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProtectedBranchesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProtectedBranchesResponseBody) *UpdateProtectedBranchesResponse
	GetBody() *UpdateProtectedBranchesResponseBody
}

type UpdateProtectedBranchesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProtectedBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProtectedBranchesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponse) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProtectedBranchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProtectedBranchesResponse) GetBody() *UpdateProtectedBranchesResponseBody {
	return s.Body
}

func (s *UpdateProtectedBranchesResponse) SetHeaders(v map[string]*string) *UpdateProtectedBranchesResponse {
	s.Headers = v
	return s
}

func (s *UpdateProtectedBranchesResponse) SetStatusCode(v int32) *UpdateProtectedBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProtectedBranchesResponse) SetBody(v *UpdateProtectedBranchesResponseBody) *UpdateProtectedBranchesResponse {
	s.Body = v
	return s
}

func (s *UpdateProtectedBranchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
