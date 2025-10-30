// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApprovalProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApprovalProcessResponse
	GetStatusCode() *int32
	SetBody(v *GetApprovalProcessResponseBody) *GetApprovalProcessResponse
	GetBody() *GetApprovalProcessResponseBody
}

type GetApprovalProcessResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApprovalProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApprovalProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponse) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApprovalProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApprovalProcessResponse) GetBody() *GetApprovalProcessResponseBody {
	return s.Body
}

func (s *GetApprovalProcessResponse) SetHeaders(v map[string]*string) *GetApprovalProcessResponse {
	s.Headers = v
	return s
}

func (s *GetApprovalProcessResponse) SetStatusCode(v int32) *GetApprovalProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApprovalProcessResponse) SetBody(v *GetApprovalProcessResponseBody) *GetApprovalProcessResponse {
	s.Body = v
	return s
}

func (s *GetApprovalProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
