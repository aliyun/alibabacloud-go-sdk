// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentStorageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentStorageResponseBody) *DeleteAgentStorageResponse
	GetBody() *DeleteAgentStorageResponseBody
}

type DeleteAgentStorageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentStorageResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentStorageResponse) GetBody() *DeleteAgentStorageResponseBody {
	return s.Body
}

func (s *DeleteAgentStorageResponse) SetHeaders(v map[string]*string) *DeleteAgentStorageResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentStorageResponse) SetStatusCode(v int32) *DeleteAgentStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentStorageResponse) SetBody(v *DeleteAgentStorageResponseBody) *DeleteAgentStorageResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
