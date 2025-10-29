// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnapshotRepoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSnapshotRepoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSnapshotRepoResponse
	GetStatusCode() *int32
	SetBody(v *AddSnapshotRepoResponseBody) *AddSnapshotRepoResponse
	GetBody() *AddSnapshotRepoResponseBody
}

type AddSnapshotRepoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSnapshotRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSnapshotRepoResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSnapshotRepoResponse) GoString() string {
	return s.String()
}

func (s *AddSnapshotRepoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSnapshotRepoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSnapshotRepoResponse) GetBody() *AddSnapshotRepoResponseBody {
	return s.Body
}

func (s *AddSnapshotRepoResponse) SetHeaders(v map[string]*string) *AddSnapshotRepoResponse {
	s.Headers = v
	return s
}

func (s *AddSnapshotRepoResponse) SetStatusCode(v int32) *AddSnapshotRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSnapshotRepoResponse) SetBody(v *AddSnapshotRepoResponseBody) *AddSnapshotRepoResponse {
	s.Body = v
	return s
}

func (s *AddSnapshotRepoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
