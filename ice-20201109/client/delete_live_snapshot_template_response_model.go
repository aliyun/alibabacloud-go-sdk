// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveSnapshotTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveSnapshotTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveSnapshotTemplateResponseBody) *DeleteLiveSnapshotTemplateResponse
	GetBody() *DeleteLiveSnapshotTemplateResponseBody
}

type DeleteLiveSnapshotTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveSnapshotTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveSnapshotTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveSnapshotTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveSnapshotTemplateResponse) GetBody() *DeleteLiveSnapshotTemplateResponseBody {
	return s.Body
}

func (s *DeleteLiveSnapshotTemplateResponse) SetHeaders(v map[string]*string) *DeleteLiveSnapshotTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveSnapshotTemplateResponse) SetStatusCode(v int32) *DeleteLiveSnapshotTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveSnapshotTemplateResponse) SetBody(v *DeleteLiveSnapshotTemplateResponseBody) *DeleteLiveSnapshotTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveSnapshotTemplateResponse) Validate() error {
	return dara.Validate(s)
}
