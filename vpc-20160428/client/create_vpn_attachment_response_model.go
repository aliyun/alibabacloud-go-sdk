// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpnAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpnAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpnAttachmentResponseBody) *CreateVpnAttachmentResponse
	GetBody() *CreateVpnAttachmentResponseBody
}

type CreateVpnAttachmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpnAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpnAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpnAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpnAttachmentResponse) GetBody() *CreateVpnAttachmentResponseBody {
	return s.Body
}

func (s *CreateVpnAttachmentResponse) SetHeaders(v map[string]*string) *CreateVpnAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateVpnAttachmentResponse) SetStatusCode(v int32) *CreateVpnAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpnAttachmentResponse) SetBody(v *CreateVpnAttachmentResponseBody) *CreateVpnAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateVpnAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
