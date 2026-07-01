// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticDBBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticDBBranchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticDBBranchResponseBody) *DeleteAgenticDBBranchResponse
	GetBody() *DeleteAgenticDBBranchResponseBody
}

type DeleteAgenticDBBranchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticDBBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticDBBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticDBBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticDBBranchResponse) GetBody() *DeleteAgenticDBBranchResponseBody {
	return s.Body
}

func (s *DeleteAgenticDBBranchResponse) SetHeaders(v map[string]*string) *DeleteAgenticDBBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticDBBranchResponse) SetStatusCode(v int32) *DeleteAgenticDBBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticDBBranchResponse) SetBody(v *DeleteAgenticDBBranchResponseBody) *DeleteAgenticDBBranchResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticDBBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
