// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *TableSnapshot) *GetTableSnapshotResponse
	GetBody() *TableSnapshot
}

type GetTableSnapshotResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TableSnapshot     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableSnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetTableSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableSnapshotResponse) GetBody() *TableSnapshot {
	return s.Body
}

func (s *GetTableSnapshotResponse) SetHeaders(v map[string]*string) *GetTableSnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetTableSnapshotResponse) SetStatusCode(v int32) *GetTableSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableSnapshotResponse) SetBody(v *TableSnapshot) *GetTableSnapshotResponse {
	s.Body = v
	return s
}

func (s *GetTableSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
