// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiskReplicaPairResponseBody) *CreateDiskReplicaPairResponse
	GetBody() *CreateDiskReplicaPairResponseBody
}

type CreateDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiskReplicaPairResponse) GetBody() *CreateDiskReplicaPairResponseBody {
	return s.Body
}

func (s *CreateDiskReplicaPairResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetStatusCode(v int32) *CreateDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetBody(v *CreateDiskReplicaPairResponseBody) *CreateDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *CreateDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
