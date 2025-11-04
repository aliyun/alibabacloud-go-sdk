// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotUrlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSnapshotUrlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSnapshotUrlsResponse
	GetStatusCode() *int32
	SetBody(v *GetSnapshotUrlsResponseBody) *GetSnapshotUrlsResponse
	GetBody() *GetSnapshotUrlsResponseBody
}

type GetSnapshotUrlsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSnapshotUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSnapshotUrlsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotUrlsResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotUrlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSnapshotUrlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSnapshotUrlsResponse) GetBody() *GetSnapshotUrlsResponseBody {
	return s.Body
}

func (s *GetSnapshotUrlsResponse) SetHeaders(v map[string]*string) *GetSnapshotUrlsResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotUrlsResponse) SetStatusCode(v int32) *GetSnapshotUrlsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSnapshotUrlsResponse) SetBody(v *GetSnapshotUrlsResponseBody) *GetSnapshotUrlsResponse {
	s.Body = v
	return s
}

func (s *GetSnapshotUrlsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
