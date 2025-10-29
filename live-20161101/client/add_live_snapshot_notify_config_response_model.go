// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveSnapshotNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveSnapshotNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveSnapshotNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveSnapshotNotifyConfigResponseBody) *AddLiveSnapshotNotifyConfigResponse
	GetBody() *AddLiveSnapshotNotifyConfigResponseBody
}

type AddLiveSnapshotNotifyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveSnapshotNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveSnapshotNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveSnapshotNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveSnapshotNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveSnapshotNotifyConfigResponse) GetBody() *AddLiveSnapshotNotifyConfigResponseBody {
	return s.Body
}

func (s *AddLiveSnapshotNotifyConfigResponse) SetHeaders(v map[string]*string) *AddLiveSnapshotNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveSnapshotNotifyConfigResponse) SetStatusCode(v int32) *AddLiveSnapshotNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigResponse) SetBody(v *AddLiveSnapshotNotifyConfigResponseBody) *AddLiveSnapshotNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveSnapshotNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
