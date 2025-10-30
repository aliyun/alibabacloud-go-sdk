// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApprovalResponse
	GetStatusCode() *int32
	SetBody(v *GetApprovalResponseBody) *GetApprovalResponse
	GetBody() *GetApprovalResponseBody
}

type GetApprovalResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponse) GoString() string {
	return s.String()
}

func (s *GetApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApprovalResponse) GetBody() *GetApprovalResponseBody {
	return s.Body
}

func (s *GetApprovalResponse) SetHeaders(v map[string]*string) *GetApprovalResponse {
	s.Headers = v
	return s
}

func (s *GetApprovalResponse) SetStatusCode(v int32) *GetApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApprovalResponse) SetBody(v *GetApprovalResponseBody) *GetApprovalResponse {
	s.Body = v
	return s
}

func (s *GetApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
