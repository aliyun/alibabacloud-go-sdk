// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveSnapshotTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveSnapshotTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveSnapshotTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveSnapshotTemplateResponseBody) *CreateLiveSnapshotTemplateResponse
	GetBody() *CreateLiveSnapshotTemplateResponseBody
}

type CreateLiveSnapshotTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveSnapshotTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveSnapshotTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveSnapshotTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveSnapshotTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveSnapshotTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveSnapshotTemplateResponse) GetBody() *CreateLiveSnapshotTemplateResponseBody {
	return s.Body
}

func (s *CreateLiveSnapshotTemplateResponse) SetHeaders(v map[string]*string) *CreateLiveSnapshotTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveSnapshotTemplateResponse) SetStatusCode(v int32) *CreateLiveSnapshotTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveSnapshotTemplateResponse) SetBody(v *CreateLiveSnapshotTemplateResponseBody) *CreateLiveSnapshotTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateLiveSnapshotTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
