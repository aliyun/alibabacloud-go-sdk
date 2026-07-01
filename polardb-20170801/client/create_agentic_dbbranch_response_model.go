// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgenticDBBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgenticDBBranchResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgenticDBBranchResponseBody) *CreateAgenticDBBranchResponse
	GetBody() *CreateAgenticDBBranchResponseBody
}

type CreateAgenticDBBranchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgenticDBBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgenticDBBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBBranchResponse) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgenticDBBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgenticDBBranchResponse) GetBody() *CreateAgenticDBBranchResponseBody {
	return s.Body
}

func (s *CreateAgenticDBBranchResponse) SetHeaders(v map[string]*string) *CreateAgenticDBBranchResponse {
	s.Headers = v
	return s
}

func (s *CreateAgenticDBBranchResponse) SetStatusCode(v int32) *CreateAgenticDBBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgenticDBBranchResponse) SetBody(v *CreateAgenticDBBranchResponseBody) *CreateAgenticDBBranchResponse {
	s.Body = v
	return s
}

func (s *CreateAgenticDBBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
