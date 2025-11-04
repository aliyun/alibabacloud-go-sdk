// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *GetSnapshotJobResponseBody) *GetSnapshotJobResponse
	GetBody() *GetSnapshotJobResponseBody
}

type GetSnapshotJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSnapshotJobResponse) GetBody() *GetSnapshotJobResponseBody {
	return s.Body
}

func (s *GetSnapshotJobResponse) SetHeaders(v map[string]*string) *GetSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotJobResponse) SetStatusCode(v int32) *GetSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSnapshotJobResponse) SetBody(v *GetSnapshotJobResponseBody) *GetSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *GetSnapshotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
