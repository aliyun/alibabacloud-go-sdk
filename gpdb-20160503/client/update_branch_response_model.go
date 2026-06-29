// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBranchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBranchResponseBody) *UpdateBranchResponse
	GetBody() *UpdateBranchResponseBody
}

type UpdateBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBranchResponse) GoString() string {
	return s.String()
}

func (s *UpdateBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBranchResponse) GetBody() *UpdateBranchResponseBody {
	return s.Body
}

func (s *UpdateBranchResponse) SetHeaders(v map[string]*string) *UpdateBranchResponse {
	s.Headers = v
	return s
}

func (s *UpdateBranchResponse) SetStatusCode(v int32) *UpdateBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBranchResponse) SetBody(v *UpdateBranchResponseBody) *UpdateBranchResponse {
	s.Body = v
	return s
}

func (s *UpdateBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
