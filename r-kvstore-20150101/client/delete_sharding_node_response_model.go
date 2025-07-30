// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteShardingNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteShardingNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteShardingNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteShardingNodeResponseBody) *DeleteShardingNodeResponse
	GetBody() *DeleteShardingNodeResponseBody
}

type DeleteShardingNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteShardingNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteShardingNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteShardingNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteShardingNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteShardingNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteShardingNodeResponse) GetBody() *DeleteShardingNodeResponseBody {
	return s.Body
}

func (s *DeleteShardingNodeResponse) SetHeaders(v map[string]*string) *DeleteShardingNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteShardingNodeResponse) SetStatusCode(v int32) *DeleteShardingNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteShardingNodeResponse) SetBody(v *DeleteShardingNodeResponseBody) *DeleteShardingNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteShardingNodeResponse) Validate() error {
	return dara.Validate(s)
}
