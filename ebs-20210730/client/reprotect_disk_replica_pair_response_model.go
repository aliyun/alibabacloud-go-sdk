// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReprotectDiskReplicaPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReprotectDiskReplicaPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReprotectDiskReplicaPairResponse
	GetStatusCode() *int32
	SetBody(v *ReprotectDiskReplicaPairResponseBody) *ReprotectDiskReplicaPairResponse
	GetBody() *ReprotectDiskReplicaPairResponseBody
}

type ReprotectDiskReplicaPairResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReprotectDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReprotectDiskReplicaPairResponse) String() string {
	return dara.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReprotectDiskReplicaPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReprotectDiskReplicaPairResponse) GetBody() *ReprotectDiskReplicaPairResponseBody {
	return s.Body
}

func (s *ReprotectDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetStatusCode(v int32) *ReprotectDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetBody(v *ReprotectDiskReplicaPairResponseBody) *ReprotectDiskReplicaPairResponse {
	s.Body = v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
