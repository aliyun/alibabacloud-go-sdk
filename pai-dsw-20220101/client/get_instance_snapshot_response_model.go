// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceSnapshotResponseBody) *GetInstanceSnapshotResponse
	GetBody() *GetInstanceSnapshotResponseBody
}

type GetInstanceSnapshotResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceSnapshotResponse) GetBody() *GetInstanceSnapshotResponseBody {
	return s.Body
}

func (s *GetInstanceSnapshotResponse) SetHeaders(v map[string]*string) *GetInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSnapshotResponse) SetStatusCode(v int32) *GetInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSnapshotResponse) SetBody(v *GetInstanceSnapshotResponseBody) *GetInstanceSnapshotResponse {
	s.Body = v
	return s
}

func (s *GetInstanceSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
