// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSnapshotTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoSnapshotTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoSnapshotTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoSnapshotTaskStatusResponseBody) *GetVideoSnapshotTaskStatusResponse
	GetBody() *GetVideoSnapshotTaskStatusResponseBody
}

type GetVideoSnapshotTaskStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoSnapshotTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoSnapshotTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSnapshotTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVideoSnapshotTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoSnapshotTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoSnapshotTaskStatusResponse) GetBody() *GetVideoSnapshotTaskStatusResponseBody {
	return s.Body
}

func (s *GetVideoSnapshotTaskStatusResponse) SetHeaders(v map[string]*string) *GetVideoSnapshotTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponse) SetStatusCode(v int32) *GetVideoSnapshotTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponse) SetBody(v *GetVideoSnapshotTaskStatusResponseBody) *GetVideoSnapshotTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
