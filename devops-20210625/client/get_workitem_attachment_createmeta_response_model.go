// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemAttachmentCreatemetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkitemAttachmentCreatemetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkitemAttachmentCreatemetaResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkitemAttachmentCreatemetaResponseBody) *GetWorkitemAttachmentCreatemetaResponse
	GetBody() *GetWorkitemAttachmentCreatemetaResponseBody
}

type GetWorkitemAttachmentCreatemetaResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkitemAttachmentCreatemetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkitemAttachmentCreatemetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemAttachmentCreatemetaResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemAttachmentCreatemetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkitemAttachmentCreatemetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkitemAttachmentCreatemetaResponse) GetBody() *GetWorkitemAttachmentCreatemetaResponseBody {
	return s.Body
}

func (s *GetWorkitemAttachmentCreatemetaResponse) SetHeaders(v map[string]*string) *GetWorkitemAttachmentCreatemetaResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponse) SetStatusCode(v int32) *GetWorkitemAttachmentCreatemetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponse) SetBody(v *GetWorkitemAttachmentCreatemetaResponseBody) *GetWorkitemAttachmentCreatemetaResponse {
	s.Body = v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
