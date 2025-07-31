// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueEngineVersionByEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueEngineVersionByEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueEngineVersionByEnvResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueEngineVersionByEnvResponseBody) *GetQueueEngineVersionByEnvResponse
	GetBody() *GetQueueEngineVersionByEnvResponseBody
}

type GetQueueEngineVersionByEnvResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueEngineVersionByEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueEngineVersionByEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueEngineVersionByEnvResponse) GoString() string {
	return s.String()
}

func (s *GetQueueEngineVersionByEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueEngineVersionByEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueEngineVersionByEnvResponse) GetBody() *GetQueueEngineVersionByEnvResponseBody {
	return s.Body
}

func (s *GetQueueEngineVersionByEnvResponse) SetHeaders(v map[string]*string) *GetQueueEngineVersionByEnvResponse {
	s.Headers = v
	return s
}

func (s *GetQueueEngineVersionByEnvResponse) SetStatusCode(v int32) *GetQueueEngineVersionByEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueEngineVersionByEnvResponse) SetBody(v *GetQueueEngineVersionByEnvResponseBody) *GetQueueEngineVersionByEnvResponse {
	s.Body = v
	return s
}

func (s *GetQueueEngineVersionByEnvResponse) Validate() error {
	return dara.Validate(s)
}
