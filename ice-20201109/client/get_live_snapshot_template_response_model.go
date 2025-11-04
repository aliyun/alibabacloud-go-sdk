// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveSnapshotTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveSnapshotTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveSnapshotTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveSnapshotTemplateResponseBody) *GetLiveSnapshotTemplateResponse
	GetBody() *GetLiveSnapshotTemplateResponseBody
}

type GetLiveSnapshotTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveSnapshotTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveSnapshotTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveSnapshotTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveSnapshotTemplateResponse) GetBody() *GetLiveSnapshotTemplateResponseBody {
	return s.Body
}

func (s *GetLiveSnapshotTemplateResponse) SetHeaders(v map[string]*string) *GetLiveSnapshotTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetLiveSnapshotTemplateResponse) SetStatusCode(v int32) *GetLiveSnapshotTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponse) SetBody(v *GetLiveSnapshotTemplateResponseBody) *GetLiveSnapshotTemplateResponse {
	s.Body = v
	return s
}

func (s *GetLiveSnapshotTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
