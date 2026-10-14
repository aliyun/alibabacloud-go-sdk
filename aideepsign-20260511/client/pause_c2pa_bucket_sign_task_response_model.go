// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseC2paBucketSignTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseC2paBucketSignTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseC2paBucketSignTaskResponse
	GetStatusCode() *int32
	SetBody(v *PauseC2paBucketSignTaskResponseBody) *PauseC2paBucketSignTaskResponse
	GetBody() *PauseC2paBucketSignTaskResponseBody
}

type PauseC2paBucketSignTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseC2paBucketSignTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseC2paBucketSignTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseC2paBucketSignTaskResponse) GoString() string {
	return s.String()
}

func (s *PauseC2paBucketSignTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseC2paBucketSignTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseC2paBucketSignTaskResponse) GetBody() *PauseC2paBucketSignTaskResponseBody {
	return s.Body
}

func (s *PauseC2paBucketSignTaskResponse) SetHeaders(v map[string]*string) *PauseC2paBucketSignTaskResponse {
	s.Headers = v
	return s
}

func (s *PauseC2paBucketSignTaskResponse) SetStatusCode(v int32) *PauseC2paBucketSignTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseC2paBucketSignTaskResponse) SetBody(v *PauseC2paBucketSignTaskResponseBody) *PauseC2paBucketSignTaskResponse {
	s.Body = v
	return s
}

func (s *PauseC2paBucketSignTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
