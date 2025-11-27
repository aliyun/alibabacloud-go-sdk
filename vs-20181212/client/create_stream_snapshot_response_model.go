// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamSnapshotResponseBody) *CreateStreamSnapshotResponse
	GetBody() *CreateStreamSnapshotResponseBody
}

type CreateStreamSnapshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamSnapshotResponse) GetBody() *CreateStreamSnapshotResponseBody {
	return s.Body
}

func (s *CreateStreamSnapshotResponse) SetHeaders(v map[string]*string) *CreateStreamSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamSnapshotResponse) SetStatusCode(v int32) *CreateStreamSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamSnapshotResponse) SetBody(v *CreateStreamSnapshotResponseBody) *CreateStreamSnapshotResponse {
	s.Body = v
	return s
}

func (s *CreateStreamSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
