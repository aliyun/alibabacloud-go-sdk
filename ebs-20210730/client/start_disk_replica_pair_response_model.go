// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *StartDiskReplicaPairResponseBody) *StartDiskReplicaPairResponse
	GetBody() *StartDiskReplicaPairResponseBody
}

type StartDiskReplicaPairResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDiskReplicaPairResponse) GetBody() *StartDiskReplicaPairResponseBody {
	return s.Body
}

func (s *StartDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StartDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaPairResponse) SetStatusCode(v int32) *StartDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiskReplicaPairResponse) SetBody(v *StartDiskReplicaPairResponseBody) *StartDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *StartDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
