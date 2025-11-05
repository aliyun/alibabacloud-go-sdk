// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDiskReplicaPairResponseBody) *RemoveDiskReplicaPairResponse
	GetBody() *RemoveDiskReplicaPairResponseBody
}

type RemoveDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDiskReplicaPairResponse) GetBody() *RemoveDiskReplicaPairResponseBody {
	return s.Body
}

func (s *RemoveDiskReplicaPairResponse) SetHeaders(v map[string]*string) *RemoveDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *RemoveDiskReplicaPairResponse) SetStatusCode(v int32) *RemoveDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDiskReplicaPairResponse) SetBody(v *RemoveDiskReplicaPairResponseBody) *RemoveDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *RemoveDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
