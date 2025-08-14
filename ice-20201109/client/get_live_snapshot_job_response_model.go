// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveSnapshotJobResponseBody) *GetLiveSnapshotJobResponse
	GetBody() *GetLiveSnapshotJobResponseBody
}

type GetLiveSnapshotJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveSnapshotJobResponse) GetBody() *GetLiveSnapshotJobResponseBody {
	return s.Body
}

func (s *GetLiveSnapshotJobResponse) SetHeaders(v map[string]*string) *GetLiveSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *GetLiveSnapshotJobResponse) SetStatusCode(v int32) *GetLiveSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveSnapshotJobResponse) SetBody(v *GetLiveSnapshotJobResponseBody) *GetLiveSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *GetLiveSnapshotJobResponse) Validate() error {
	return dara.Validate(s)
}
