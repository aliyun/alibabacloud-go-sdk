// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemArtifactAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemArtifactAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemArtifactAccessResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemArtifactAccessResponseBody) *GetWorkItemArtifactAccessResponse
	GetBody() *GetWorkItemArtifactAccessResponseBody
}

type GetWorkItemArtifactAccessResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemArtifactAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemArtifactAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemArtifactAccessResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemArtifactAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemArtifactAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemArtifactAccessResponse) GetBody() *GetWorkItemArtifactAccessResponseBody {
	return s.Body
}

func (s *GetWorkItemArtifactAccessResponse) SetHeaders(v map[string]*string) *GetWorkItemArtifactAccessResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemArtifactAccessResponse) SetStatusCode(v int32) *GetWorkItemArtifactAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemArtifactAccessResponse) SetBody(v *GetWorkItemArtifactAccessResponseBody) *GetWorkItemArtifactAccessResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemArtifactAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
