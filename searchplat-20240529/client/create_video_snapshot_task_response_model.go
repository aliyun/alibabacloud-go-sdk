// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSnapshotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoSnapshotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoSnapshotTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoSnapshotTaskResponseBody) *CreateVideoSnapshotTaskResponse
	GetBody() *CreateVideoSnapshotTaskResponseBody
}

type CreateVideoSnapshotTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoSnapshotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoSnapshotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSnapshotTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoSnapshotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoSnapshotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoSnapshotTaskResponse) GetBody() *CreateVideoSnapshotTaskResponseBody {
	return s.Body
}

func (s *CreateVideoSnapshotTaskResponse) SetHeaders(v map[string]*string) *CreateVideoSnapshotTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoSnapshotTaskResponse) SetStatusCode(v int32) *CreateVideoSnapshotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoSnapshotTaskResponse) SetBody(v *CreateVideoSnapshotTaskResponseBody) *CreateVideoSnapshotTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoSnapshotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
