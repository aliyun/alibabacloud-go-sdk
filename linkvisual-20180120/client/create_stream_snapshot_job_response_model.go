// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamSnapshotJobResponseBody) *CreateStreamSnapshotJobResponse
	GetBody() *CreateStreamSnapshotJobResponseBody
}

type CreateStreamSnapshotJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamSnapshotJobResponse) GetBody() *CreateStreamSnapshotJobResponseBody {
	return s.Body
}

func (s *CreateStreamSnapshotJobResponse) SetHeaders(v map[string]*string) *CreateStreamSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamSnapshotJobResponse) SetStatusCode(v int32) *CreateStreamSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamSnapshotJobResponse) SetBody(v *CreateStreamSnapshotJobResponseBody) *CreateStreamSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *CreateStreamSnapshotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
