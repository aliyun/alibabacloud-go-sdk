// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAppSnapshotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveAppSnapshotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveAppSnapshotConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveAppSnapshotConfigResponseBody) *UpdateLiveAppSnapshotConfigResponse
	GetBody() *UpdateLiveAppSnapshotConfigResponseBody
}

type UpdateLiveAppSnapshotConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveAppSnapshotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveAppSnapshotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppSnapshotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveAppSnapshotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveAppSnapshotConfigResponse) GetBody() *UpdateLiveAppSnapshotConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveAppSnapshotConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveAppSnapshotConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveAppSnapshotConfigResponse) SetStatusCode(v int32) *UpdateLiveAppSnapshotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigResponse) SetBody(v *UpdateLiveAppSnapshotConfigResponseBody) *UpdateLiveAppSnapshotConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveAppSnapshotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
