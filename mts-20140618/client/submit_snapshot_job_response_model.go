// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSnapshotJobResponseBody) *SubmitSnapshotJobResponse
	GetBody() *SubmitSnapshotJobResponseBody
}

type SubmitSnapshotJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSnapshotJobResponse) GetBody() *SubmitSnapshotJobResponseBody {
	return s.Body
}

func (s *SubmitSnapshotJobResponse) SetHeaders(v map[string]*string) *SubmitSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSnapshotJobResponse) SetStatusCode(v int32) *SubmitSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSnapshotJobResponse) SetBody(v *SubmitSnapshotJobResponseBody) *SubmitSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSnapshotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
