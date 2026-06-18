// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBranchResponse
	GetStatusCode() *int32
	SetBody(v *CreateBranchResponseBody) *CreateBranchResponse
	GetBody() *CreateBranchResponseBody
}

type CreateBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchResponse) GoString() string {
	return s.String()
}

func (s *CreateBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBranchResponse) GetBody() *CreateBranchResponseBody {
	return s.Body
}

func (s *CreateBranchResponse) SetHeaders(v map[string]*string) *CreateBranchResponse {
	s.Headers = v
	return s
}

func (s *CreateBranchResponse) SetStatusCode(v int32) *CreateBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBranchResponse) SetBody(v *CreateBranchResponseBody) *CreateBranchResponse {
	s.Body = v
	return s
}

func (s *CreateBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
