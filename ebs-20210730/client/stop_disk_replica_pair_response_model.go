// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *StopDiskReplicaPairResponseBody) *StopDiskReplicaPairResponse
	GetBody() *StopDiskReplicaPairResponseBody
}

type StopDiskReplicaPairResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDiskReplicaPairResponse) GetBody() *StopDiskReplicaPairResponseBody {
	return s.Body
}

func (s *StopDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StopDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaPairResponse) SetStatusCode(v int32) *StopDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDiskReplicaPairResponse) SetBody(v *StopDiskReplicaPairResponseBody) *StopDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *StopDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
