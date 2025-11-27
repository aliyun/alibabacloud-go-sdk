// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCSnapshotResponseBody) *CreateRCSnapshotResponse
	GetBody() *CreateRCSnapshotResponseBody
}

type CreateRCSnapshotResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateRCSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCSnapshotResponse) GetBody() *CreateRCSnapshotResponseBody {
	return s.Body
}

func (s *CreateRCSnapshotResponse) SetHeaders(v map[string]*string) *CreateRCSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateRCSnapshotResponse) SetStatusCode(v int32) *CreateRCSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCSnapshotResponse) SetBody(v *CreateRCSnapshotResponseBody) *CreateRCSnapshotResponse {
	s.Body = v
	return s
}

func (s *CreateRCSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
