// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveSnapshotJobCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendLiveSnapshotJobCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendLiveSnapshotJobCommandResponse
	GetStatusCode() *int32
	SetBody(v *SendLiveSnapshotJobCommandResponseBody) *SendLiveSnapshotJobCommandResponse
	GetBody() *SendLiveSnapshotJobCommandResponseBody
}

type SendLiveSnapshotJobCommandResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLiveSnapshotJobCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLiveSnapshotJobCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s SendLiveSnapshotJobCommandResponse) GoString() string {
	return s.String()
}

func (s *SendLiveSnapshotJobCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendLiveSnapshotJobCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendLiveSnapshotJobCommandResponse) GetBody() *SendLiveSnapshotJobCommandResponseBody {
	return s.Body
}

func (s *SendLiveSnapshotJobCommandResponse) SetHeaders(v map[string]*string) *SendLiveSnapshotJobCommandResponse {
	s.Headers = v
	return s
}

func (s *SendLiveSnapshotJobCommandResponse) SetStatusCode(v int32) *SendLiveSnapshotJobCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLiveSnapshotJobCommandResponse) SetBody(v *SendLiveSnapshotJobCommandResponseBody) *SendLiveSnapshotJobCommandResponse {
	s.Body = v
	return s
}

func (s *SendLiveSnapshotJobCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
