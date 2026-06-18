// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBranchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBranchResponseBody) *DeleteBranchResponse
	GetBody() *DeleteBranchResponseBody
}

type DeleteBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBranchResponse) GetBody() *DeleteBranchResponseBody {
	return s.Body
}

func (s *DeleteBranchResponse) SetHeaders(v map[string]*string) *DeleteBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteBranchResponse) SetStatusCode(v int32) *DeleteBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBranchResponse) SetBody(v *DeleteBranchResponseBody) *DeleteBranchResponse {
	s.Body = v
	return s
}

func (s *DeleteBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
