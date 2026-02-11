// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsistencySnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsistencySnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsistencySnapshotResponse
	GetStatusCode() *int32
	SetBody(v *GetConsistencySnapshotResponseBody) *GetConsistencySnapshotResponse
	GetBody() *GetConsistencySnapshotResponseBody
}

type GetConsistencySnapshotResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsistencySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsistencySnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsistencySnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsistencySnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsistencySnapshotResponse) GetBody() *GetConsistencySnapshotResponseBody {
	return s.Body
}

func (s *GetConsistencySnapshotResponse) SetHeaders(v map[string]*string) *GetConsistencySnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetConsistencySnapshotResponse) SetStatusCode(v int32) *GetConsistencySnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsistencySnapshotResponse) SetBody(v *GetConsistencySnapshotResponseBody) *GetConsistencySnapshotResponse {
	s.Body = v
	return s
}

func (s *GetConsistencySnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
