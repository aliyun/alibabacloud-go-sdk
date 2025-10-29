// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAppSnapshotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveAppSnapshotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveAppSnapshotConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveAppSnapshotConfigResponseBody) *AddLiveAppSnapshotConfigResponse
	GetBody() *AddLiveAppSnapshotConfigResponseBody
}

type AddLiveAppSnapshotConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveAppSnapshotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveAppSnapshotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAppSnapshotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveAppSnapshotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveAppSnapshotConfigResponse) GetBody() *AddLiveAppSnapshotConfigResponseBody {
	return s.Body
}

func (s *AddLiveAppSnapshotConfigResponse) SetHeaders(v map[string]*string) *AddLiveAppSnapshotConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveAppSnapshotConfigResponse) SetStatusCode(v int32) *AddLiveAppSnapshotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveAppSnapshotConfigResponse) SetBody(v *AddLiveAppSnapshotConfigResponseBody) *AddLiveAppSnapshotConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveAppSnapshotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
