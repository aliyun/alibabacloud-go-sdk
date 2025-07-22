// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceSnapshotResponseBody) *CreateInstanceSnapshotResponse
	GetBody() *CreateInstanceSnapshotResponseBody
}

type CreateInstanceSnapshotResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceSnapshotResponse) GetBody() *CreateInstanceSnapshotResponseBody {
	return s.Body
}

func (s *CreateInstanceSnapshotResponse) SetHeaders(v map[string]*string) *CreateInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceSnapshotResponse) SetStatusCode(v int32) *CreateInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceSnapshotResponse) SetBody(v *CreateInstanceSnapshotResponseBody) *CreateInstanceSnapshotResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
