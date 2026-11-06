// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCollaborationGroupOutputItemPreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCollaborationGroupOutputItemPreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCollaborationGroupOutputItemPreviewResponse
	GetStatusCode() *int32
	SetBody(v *GetCollaborationGroupOutputItemPreviewResponseBody) *GetCollaborationGroupOutputItemPreviewResponse
	GetBody() *GetCollaborationGroupOutputItemPreviewResponseBody
}

type GetCollaborationGroupOutputItemPreviewResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCollaborationGroupOutputItemPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCollaborationGroupOutputItemPreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCollaborationGroupOutputItemPreviewResponse) GoString() string {
	return s.String()
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) GetBody() *GetCollaborationGroupOutputItemPreviewResponseBody {
	return s.Body
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) SetHeaders(v map[string]*string) *GetCollaborationGroupOutputItemPreviewResponse {
	s.Headers = v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) SetStatusCode(v int32) *GetCollaborationGroupOutputItemPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) SetBody(v *GetCollaborationGroupOutputItemPreviewResponseBody) *GetCollaborationGroupOutputItemPreviewResponse {
	s.Body = v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
