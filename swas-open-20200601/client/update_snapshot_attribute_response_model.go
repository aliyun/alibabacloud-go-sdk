// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSnapshotAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSnapshotAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSnapshotAttributeResponseBody) *UpdateSnapshotAttributeResponse
	GetBody() *UpdateSnapshotAttributeResponseBody
}

type UpdateSnapshotAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSnapshotAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSnapshotAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSnapshotAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSnapshotAttributeResponse) GetBody() *UpdateSnapshotAttributeResponseBody {
	return s.Body
}

func (s *UpdateSnapshotAttributeResponse) SetHeaders(v map[string]*string) *UpdateSnapshotAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotAttributeResponse) SetStatusCode(v int32) *UpdateSnapshotAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSnapshotAttributeResponse) SetBody(v *UpdateSnapshotAttributeResponseBody) *UpdateSnapshotAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateSnapshotAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
