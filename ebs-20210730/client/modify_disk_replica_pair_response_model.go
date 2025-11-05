// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskReplicaPairResponseBody) *ModifyDiskReplicaPairResponse
	GetBody() *ModifyDiskReplicaPairResponseBody
}

type ModifyDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskReplicaPairResponse) GetBody() *ModifyDiskReplicaPairResponseBody {
	return s.Body
}

func (s *ModifyDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetStatusCode(v int32) *ModifyDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetBody(v *ModifyDiskReplicaPairResponseBody) *ModifyDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
