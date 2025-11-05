// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *AddDiskReplicaPairResponseBody) *AddDiskReplicaPairResponse
	GetBody() *AddDiskReplicaPairResponseBody
}

type AddDiskReplicaPairResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDiskReplicaPairResponse) GetBody() *AddDiskReplicaPairResponseBody {
	return s.Body
}

func (s *AddDiskReplicaPairResponse) SetHeaders(v map[string]*string) *AddDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *AddDiskReplicaPairResponse) SetStatusCode(v int32) *AddDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDiskReplicaPairResponse) SetBody(v *AddDiskReplicaPairResponseBody) *AddDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *AddDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
