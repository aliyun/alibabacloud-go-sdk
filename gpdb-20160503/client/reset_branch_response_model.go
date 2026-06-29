// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetBranchResponse
	GetStatusCode() *int32
	SetBody(v *ResetBranchResponseBody) *ResetBranchResponse
	GetBody() *ResetBranchResponseBody
}

type ResetBranchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetBranchResponse) GoString() string {
	return s.String()
}

func (s *ResetBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetBranchResponse) GetBody() *ResetBranchResponseBody {
	return s.Body
}

func (s *ResetBranchResponse) SetHeaders(v map[string]*string) *ResetBranchResponse {
	s.Headers = v
	return s
}

func (s *ResetBranchResponse) SetStatusCode(v int32) *ResetBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetBranchResponse) SetBody(v *ResetBranchResponseBody) *ResetBranchResponse {
	s.Body = v
	return s
}

func (s *ResetBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
