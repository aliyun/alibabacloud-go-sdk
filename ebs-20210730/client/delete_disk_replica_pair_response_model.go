// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiskReplicaPairResponseBody) *DeleteDiskReplicaPairResponse
	GetBody() *DeleteDiskReplicaPairResponseBody
}

type DeleteDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiskReplicaPairResponse) GetBody() *DeleteDiskReplicaPairResponseBody {
	return s.Body
}

func (s *DeleteDiskReplicaPairResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetStatusCode(v int32) *DeleteDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetBody(v *DeleteDiskReplicaPairResponseBody) *DeleteDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *DeleteDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
