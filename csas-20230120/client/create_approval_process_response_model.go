// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApprovalProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApprovalProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApprovalProcessResponse
	GetStatusCode() *int32
	SetBody(v *CreateApprovalProcessResponseBody) *CreateApprovalProcessResponse
	GetBody() *CreateApprovalProcessResponseBody
}

type CreateApprovalProcessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApprovalProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApprovalProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApprovalProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApprovalProcessResponse) GetBody() *CreateApprovalProcessResponseBody {
	return s.Body
}

func (s *CreateApprovalProcessResponse) SetHeaders(v map[string]*string) *CreateApprovalProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateApprovalProcessResponse) SetStatusCode(v int32) *CreateApprovalProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApprovalProcessResponse) SetBody(v *CreateApprovalProcessResponseBody) *CreateApprovalProcessResponse {
	s.Body = v
	return s
}

func (s *CreateApprovalProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
