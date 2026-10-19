// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAttachmentAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemAttachmentAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemAttachmentAccessResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemAttachmentAccessResponseBody) *GetWorkItemAttachmentAccessResponse
	GetBody() *GetWorkItemAttachmentAccessResponseBody
}

type GetWorkItemAttachmentAccessResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemAttachmentAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemAttachmentAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAttachmentAccessResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemAttachmentAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemAttachmentAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemAttachmentAccessResponse) GetBody() *GetWorkItemAttachmentAccessResponseBody {
	return s.Body
}

func (s *GetWorkItemAttachmentAccessResponse) SetHeaders(v map[string]*string) *GetWorkItemAttachmentAccessResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemAttachmentAccessResponse) SetStatusCode(v int32) *GetWorkItemAttachmentAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemAttachmentAccessResponse) SetBody(v *GetWorkItemAttachmentAccessResponseBody) *GetWorkItemAttachmentAccessResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemAttachmentAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
