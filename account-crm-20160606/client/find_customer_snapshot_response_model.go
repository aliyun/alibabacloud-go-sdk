// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomerSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindCustomerSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindCustomerSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *FindCustomerSnapshotResponseBody) *FindCustomerSnapshotResponse
	GetBody() *FindCustomerSnapshotResponseBody
}

type FindCustomerSnapshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindCustomerSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindCustomerSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerSnapshotResponse) GoString() string {
	return s.String()
}

func (s *FindCustomerSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindCustomerSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindCustomerSnapshotResponse) GetBody() *FindCustomerSnapshotResponseBody {
	return s.Body
}

func (s *FindCustomerSnapshotResponse) SetHeaders(v map[string]*string) *FindCustomerSnapshotResponse {
	s.Headers = v
	return s
}

func (s *FindCustomerSnapshotResponse) SetStatusCode(v int32) *FindCustomerSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *FindCustomerSnapshotResponse) SetBody(v *FindCustomerSnapshotResponseBody) *FindCustomerSnapshotResponse {
	s.Body = v
	return s
}

func (s *FindCustomerSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
