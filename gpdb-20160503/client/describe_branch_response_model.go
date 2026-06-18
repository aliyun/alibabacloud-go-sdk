// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBranchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBranchResponseBody) *DescribeBranchResponse
	GetBody() *DescribeBranchResponseBody
}

type DescribeBranchResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBranchResponse) GoString() string {
	return s.String()
}

func (s *DescribeBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBranchResponse) GetBody() *DescribeBranchResponseBody {
	return s.Body
}

func (s *DescribeBranchResponse) SetHeaders(v map[string]*string) *DescribeBranchResponse {
	s.Headers = v
	return s
}

func (s *DescribeBranchResponse) SetStatusCode(v int32) *DescribeBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBranchResponse) SetBody(v *DescribeBranchResponseBody) *DescribeBranchResponse {
	s.Body = v
	return s
}

func (s *DescribeBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
