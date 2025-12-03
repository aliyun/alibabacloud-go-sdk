// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectdBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProtectdBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProtectdBranchResponse
	GetStatusCode() *int32
	SetBody(v *CreateProtectdBranchResponseBody) *CreateProtectdBranchResponse
	GetBody() *CreateProtectdBranchResponseBody
}

type CreateProtectdBranchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProtectdBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProtectdBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponse) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProtectdBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProtectdBranchResponse) GetBody() *CreateProtectdBranchResponseBody {
	return s.Body
}

func (s *CreateProtectdBranchResponse) SetHeaders(v map[string]*string) *CreateProtectdBranchResponse {
	s.Headers = v
	return s
}

func (s *CreateProtectdBranchResponse) SetStatusCode(v int32) *CreateProtectdBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtectdBranchResponse) SetBody(v *CreateProtectdBranchResponseBody) *CreateProtectdBranchResponse {
	s.Body = v
	return s
}

func (s *CreateProtectdBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
