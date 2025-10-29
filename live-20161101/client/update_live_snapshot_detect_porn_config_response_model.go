// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotDetectPornConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveSnapshotDetectPornConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveSnapshotDetectPornConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveSnapshotDetectPornConfigResponseBody) *UpdateLiveSnapshotDetectPornConfigResponse
	GetBody() *UpdateLiveSnapshotDetectPornConfigResponseBody
}

type UpdateLiveSnapshotDetectPornConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveSnapshotDetectPornConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveSnapshotDetectPornConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) GetBody() *UpdateLiveSnapshotDetectPornConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveSnapshotDetectPornConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) SetStatusCode(v int32) *UpdateLiveSnapshotDetectPornConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) SetBody(v *UpdateLiveSnapshotDetectPornConfigResponseBody) *UpdateLiveSnapshotDetectPornConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
