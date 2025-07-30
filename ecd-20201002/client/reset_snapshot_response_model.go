// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *ResetSnapshotResponseBody) *ResetSnapshotResponse
	GetBody() *ResetSnapshotResponseBody
}

type ResetSnapshotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetSnapshotResponse) GetBody() *ResetSnapshotResponseBody {
	return s.Body
}

func (s *ResetSnapshotResponse) SetHeaders(v map[string]*string) *ResetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ResetSnapshotResponse) SetStatusCode(v int32) *ResetSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSnapshotResponse) SetBody(v *ResetSnapshotResponseBody) *ResetSnapshotResponse {
	s.Body = v
	return s
}

func (s *ResetSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
