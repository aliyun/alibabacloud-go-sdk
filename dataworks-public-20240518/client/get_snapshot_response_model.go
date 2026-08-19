// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *GetSnapshotResponseBody) *GetSnapshotResponse
	GetBody() *GetSnapshotResponseBody
}

type GetSnapshotResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSnapshotResponse) GetBody() *GetSnapshotResponseBody {
	return s.Body
}

func (s *GetSnapshotResponse) SetHeaders(v map[string]*string) *GetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotResponse) SetStatusCode(v int32) *GetSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSnapshotResponse) SetBody(v *GetSnapshotResponseBody) *GetSnapshotResponse {
	s.Body = v
	return s
}

func (s *GetSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
