// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCloudPhoneNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewCloudPhoneNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewCloudPhoneNodesResponse
	GetStatusCode() *int32
	SetBody(v *RenewCloudPhoneNodesResponseBody) *RenewCloudPhoneNodesResponse
	GetBody() *RenewCloudPhoneNodesResponseBody
}

type RenewCloudPhoneNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewCloudPhoneNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewCloudPhoneNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewCloudPhoneNodesResponse) GetBody() *RenewCloudPhoneNodesResponseBody {
	return s.Body
}

func (s *RenewCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *RenewCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *RenewCloudPhoneNodesResponse) SetStatusCode(v int32) *RenewCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewCloudPhoneNodesResponse) SetBody(v *RenewCloudPhoneNodesResponseBody) *RenewCloudPhoneNodesResponse {
	s.Body = v
	return s
}

func (s *RenewCloudPhoneNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
