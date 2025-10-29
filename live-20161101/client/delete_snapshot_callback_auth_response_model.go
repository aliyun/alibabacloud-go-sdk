// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotCallbackAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnapshotCallbackAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnapshotCallbackAuthResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnapshotCallbackAuthResponseBody) *DeleteSnapshotCallbackAuthResponse
	GetBody() *DeleteSnapshotCallbackAuthResponseBody
}

type DeleteSnapshotCallbackAuthResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotCallbackAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotCallbackAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotCallbackAuthResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotCallbackAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnapshotCallbackAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnapshotCallbackAuthResponse) GetBody() *DeleteSnapshotCallbackAuthResponseBody {
	return s.Body
}

func (s *DeleteSnapshotCallbackAuthResponse) SetHeaders(v map[string]*string) *DeleteSnapshotCallbackAuthResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotCallbackAuthResponse) SetStatusCode(v int32) *DeleteSnapshotCallbackAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotCallbackAuthResponse) SetBody(v *DeleteSnapshotCallbackAuthResponseBody) *DeleteSnapshotCallbackAuthResponse {
	s.Body = v
	return s
}

func (s *DeleteSnapshotCallbackAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
