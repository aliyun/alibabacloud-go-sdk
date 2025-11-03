// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnAttachmentAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnAttachmentAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnAttachmentAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnAttachmentAttributeResponseBody) *ModifyVpnAttachmentAttributeResponse
	GetBody() *ModifyVpnAttachmentAttributeResponseBody
}

type ModifyVpnAttachmentAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnAttachmentAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnAttachmentAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnAttachmentAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnAttachmentAttributeResponse) GetBody() *ModifyVpnAttachmentAttributeResponseBody {
	return s.Body
}

func (s *ModifyVpnAttachmentAttributeResponse) SetHeaders(v map[string]*string) *ModifyVpnAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponse) SetStatusCode(v int32) *ModifyVpnAttachmentAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponse) SetBody(v *ModifyVpnAttachmentAttributeResponseBody) *ModifyVpnAttachmentAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnAttachmentAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
