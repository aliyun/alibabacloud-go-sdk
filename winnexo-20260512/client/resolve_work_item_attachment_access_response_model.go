// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemAttachmentAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResolveWorkItemAttachmentAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResolveWorkItemAttachmentAccessResponse
	GetStatusCode() *int32
	SetBody(v *ResolveWorkItemAttachmentAccessResponseBody) *ResolveWorkItemAttachmentAccessResponse
	GetBody() *ResolveWorkItemAttachmentAccessResponseBody
}

type ResolveWorkItemAttachmentAccessResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResolveWorkItemAttachmentAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResolveWorkItemAttachmentAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemAttachmentAccessResponse) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemAttachmentAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResolveWorkItemAttachmentAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResolveWorkItemAttachmentAccessResponse) GetBody() *ResolveWorkItemAttachmentAccessResponseBody {
	return s.Body
}

func (s *ResolveWorkItemAttachmentAccessResponse) SetHeaders(v map[string]*string) *ResolveWorkItemAttachmentAccessResponse {
	s.Headers = v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponse) SetStatusCode(v int32) *ResolveWorkItemAttachmentAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponse) SetBody(v *ResolveWorkItemAttachmentAccessResponseBody) *ResolveWorkItemAttachmentAccessResponse {
	s.Body = v
	return s
}

func (s *ResolveWorkItemAttachmentAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
