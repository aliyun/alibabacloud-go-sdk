// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairKVCacheVNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTairKVCacheVNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTairKVCacheVNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateTairKVCacheVNodeResponseBody) *CreateTairKVCacheVNodeResponse
	GetBody() *CreateTairKVCacheVNodeResponseBody
}

type CreateTairKVCacheVNodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTairKVCacheVNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTairKVCacheVNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTairKVCacheVNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateTairKVCacheVNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTairKVCacheVNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTairKVCacheVNodeResponse) GetBody() *CreateTairKVCacheVNodeResponseBody {
	return s.Body
}

func (s *CreateTairKVCacheVNodeResponse) SetHeaders(v map[string]*string) *CreateTairKVCacheVNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateTairKVCacheVNodeResponse) SetStatusCode(v int32) *CreateTairKVCacheVNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponse) SetBody(v *CreateTairKVCacheVNodeResponseBody) *CreateTairKVCacheVNodeResponse {
	s.Body = v
	return s
}

func (s *CreateTairKVCacheVNodeResponse) Validate() error {
	return dara.Validate(s)
}
