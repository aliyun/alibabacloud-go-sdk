// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveSnapshotTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveSnapshotTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveSnapshotTemplateResponseBody) *UpdateLiveSnapshotTemplateResponse
	GetBody() *UpdateLiveSnapshotTemplateResponseBody
}

type UpdateLiveSnapshotTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveSnapshotTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveSnapshotTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveSnapshotTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveSnapshotTemplateResponse) GetBody() *UpdateLiveSnapshotTemplateResponseBody {
	return s.Body
}

func (s *UpdateLiveSnapshotTemplateResponse) SetHeaders(v map[string]*string) *UpdateLiveSnapshotTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveSnapshotTemplateResponse) SetStatusCode(v int32) *UpdateLiveSnapshotTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateResponse) SetBody(v *UpdateLiveSnapshotTemplateResponseBody) *UpdateLiveSnapshotTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveSnapshotTemplateResponse) Validate() error {
	return dara.Validate(s)
}
