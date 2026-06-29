// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreBranchResponse
	GetStatusCode() *int32
	SetBody(v *RestoreBranchResponseBody) *RestoreBranchResponse
	GetBody() *RestoreBranchResponseBody
}

type RestoreBranchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreBranchResponse) GoString() string {
	return s.String()
}

func (s *RestoreBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreBranchResponse) GetBody() *RestoreBranchResponseBody {
	return s.Body
}

func (s *RestoreBranchResponse) SetHeaders(v map[string]*string) *RestoreBranchResponse {
	s.Headers = v
	return s
}

func (s *RestoreBranchResponse) SetStatusCode(v int32) *RestoreBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreBranchResponse) SetBody(v *RestoreBranchResponseBody) *RestoreBranchResponse {
	s.Body = v
	return s
}

func (s *RestoreBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
