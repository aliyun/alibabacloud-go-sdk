// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FailoverDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FailoverDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *FailoverDiskReplicaPairResponseBody) *FailoverDiskReplicaPairResponse
	GetBody() *FailoverDiskReplicaPairResponseBody
}

type FailoverDiskReplicaPairResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s FailoverDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FailoverDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FailoverDiskReplicaPairResponse) GetBody() *FailoverDiskReplicaPairResponseBody {
	return s.Body
}

func (s *FailoverDiskReplicaPairResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetStatusCode(v int32) *FailoverDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetBody(v *FailoverDiskReplicaPairResponseBody) *FailoverDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *FailoverDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
