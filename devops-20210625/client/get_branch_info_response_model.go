// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBranchInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBranchInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBranchInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBranchInfoResponseBody) *GetBranchInfoResponse
	GetBody() *GetBranchInfoResponseBody
}

type GetBranchInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBranchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBranchInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBranchInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBranchInfoResponse) GetBody() *GetBranchInfoResponseBody {
	return s.Body
}

func (s *GetBranchInfoResponse) SetHeaders(v map[string]*string) *GetBranchInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBranchInfoResponse) SetStatusCode(v int32) *GetBranchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBranchInfoResponse) SetBody(v *GetBranchInfoResponseBody) *GetBranchInfoResponse {
	s.Body = v
	return s
}

func (s *GetBranchInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
