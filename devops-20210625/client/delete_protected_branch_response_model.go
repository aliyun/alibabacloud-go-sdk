// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtectedBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProtectedBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProtectedBranchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProtectedBranchResponseBody) *DeleteProtectedBranchResponse
	GetBody() *DeleteProtectedBranchResponseBody
}

type DeleteProtectedBranchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProtectedBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProtectedBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectedBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProtectedBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProtectedBranchResponse) GetBody() *DeleteProtectedBranchResponseBody {
	return s.Body
}

func (s *DeleteProtectedBranchResponse) SetHeaders(v map[string]*string) *DeleteProtectedBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtectedBranchResponse) SetStatusCode(v int32) *DeleteProtectedBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtectedBranchResponse) SetBody(v *DeleteProtectedBranchResponseBody) *DeleteProtectedBranchResponse {
	s.Body = v
	return s
}

func (s *DeleteProtectedBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
