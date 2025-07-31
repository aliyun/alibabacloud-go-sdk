// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterQueueInfoByEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterQueueInfoByEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterQueueInfoByEnvResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterQueueInfoByEnvResponseBody) *GetClusterQueueInfoByEnvResponse
	GetBody() *GetClusterQueueInfoByEnvResponseBody
}

type GetClusterQueueInfoByEnvResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterQueueInfoByEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterQueueInfoByEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterQueueInfoByEnvResponse) GoString() string {
	return s.String()
}

func (s *GetClusterQueueInfoByEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterQueueInfoByEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterQueueInfoByEnvResponse) GetBody() *GetClusterQueueInfoByEnvResponseBody {
	return s.Body
}

func (s *GetClusterQueueInfoByEnvResponse) SetHeaders(v map[string]*string) *GetClusterQueueInfoByEnvResponse {
	s.Headers = v
	return s
}

func (s *GetClusterQueueInfoByEnvResponse) SetStatusCode(v int32) *GetClusterQueueInfoByEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponse) SetBody(v *GetClusterQueueInfoByEnvResponseBody) *GetClusterQueueInfoByEnvResponse {
	s.Body = v
	return s
}

func (s *GetClusterQueueInfoByEnvResponse) Validate() error {
	return dara.Validate(s)
}
