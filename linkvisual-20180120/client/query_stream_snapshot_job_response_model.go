// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryStreamSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryStreamSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryStreamSnapshotJobResponseBody) *QueryStreamSnapshotJobResponse
	GetBody() *QueryStreamSnapshotJobResponseBody
}

type QueryStreamSnapshotJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStreamSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStreamSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryStreamSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryStreamSnapshotJobResponse) GetBody() *QueryStreamSnapshotJobResponseBody {
	return s.Body
}

func (s *QueryStreamSnapshotJobResponse) SetHeaders(v map[string]*string) *QueryStreamSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *QueryStreamSnapshotJobResponse) SetStatusCode(v int32) *QueryStreamSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStreamSnapshotJobResponse) SetBody(v *QueryStreamSnapshotJobResponseBody) *QueryStreamSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *QueryStreamSnapshotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
