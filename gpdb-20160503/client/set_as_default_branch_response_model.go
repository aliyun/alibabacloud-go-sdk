// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAsDefaultBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAsDefaultBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAsDefaultBranchResponse
	GetStatusCode() *int32
	SetBody(v *SetAsDefaultBranchResponseBody) *SetAsDefaultBranchResponse
	GetBody() *SetAsDefaultBranchResponseBody
}

type SetAsDefaultBranchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAsDefaultBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAsDefaultBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAsDefaultBranchResponse) GoString() string {
	return s.String()
}

func (s *SetAsDefaultBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAsDefaultBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAsDefaultBranchResponse) GetBody() *SetAsDefaultBranchResponseBody {
	return s.Body
}

func (s *SetAsDefaultBranchResponse) SetHeaders(v map[string]*string) *SetAsDefaultBranchResponse {
	s.Headers = v
	return s
}

func (s *SetAsDefaultBranchResponse) SetStatusCode(v int32) *SetAsDefaultBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAsDefaultBranchResponse) SetBody(v *SetAsDefaultBranchResponseBody) *SetAsDefaultBranchResponse {
	s.Body = v
	return s
}

func (s *SetAsDefaultBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
