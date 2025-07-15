// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveSnapshotDetectPornConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveSnapshotDetectPornConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveSnapshotDetectPornConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveSnapshotDetectPornConfigResponseBody) *AddLiveSnapshotDetectPornConfigResponse
	GetBody() *AddLiveSnapshotDetectPornConfigResponseBody
}

type AddLiveSnapshotDetectPornConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveSnapshotDetectPornConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveSnapshotDetectPornConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotDetectPornConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveSnapshotDetectPornConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveSnapshotDetectPornConfigResponse) GetBody() *AddLiveSnapshotDetectPornConfigResponseBody {
	return s.Body
}

func (s *AddLiveSnapshotDetectPornConfigResponse) SetHeaders(v map[string]*string) *AddLiveSnapshotDetectPornConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigResponse) SetStatusCode(v int32) *AddLiveSnapshotDetectPornConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigResponse) SetBody(v *AddLiveSnapshotDetectPornConfigResponseBody) *AddLiveSnapshotDetectPornConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigResponse) Validate() error {
	return dara.Validate(s)
}
