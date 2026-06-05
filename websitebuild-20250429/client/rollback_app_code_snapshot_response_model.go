// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppCodeSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackAppCodeSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackAppCodeSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *RollbackAppCodeSnapshotResponseBody) *RollbackAppCodeSnapshotResponse
	GetBody() *RollbackAppCodeSnapshotResponseBody
}

type RollbackAppCodeSnapshotResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackAppCodeSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackAppCodeSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppCodeSnapshotResponse) GoString() string {
	return s.String()
}

func (s *RollbackAppCodeSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackAppCodeSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackAppCodeSnapshotResponse) GetBody() *RollbackAppCodeSnapshotResponseBody {
	return s.Body
}

func (s *RollbackAppCodeSnapshotResponse) SetHeaders(v map[string]*string) *RollbackAppCodeSnapshotResponse {
	s.Headers = v
	return s
}

func (s *RollbackAppCodeSnapshotResponse) SetStatusCode(v int32) *RollbackAppCodeSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponse) SetBody(v *RollbackAppCodeSnapshotResponseBody) *RollbackAppCodeSnapshotResponse {
	s.Body = v
	return s
}

func (s *RollbackAppCodeSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
