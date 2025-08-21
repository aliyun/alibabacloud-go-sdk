// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShardRecoveriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListShardRecoveriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListShardRecoveriesResponse
	GetStatusCode() *int32
	SetBody(v *ListShardRecoveriesResponseBody) *ListShardRecoveriesResponse
	GetBody() *ListShardRecoveriesResponseBody
}

type ListShardRecoveriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShardRecoveriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShardRecoveriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListShardRecoveriesResponse) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListShardRecoveriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListShardRecoveriesResponse) GetBody() *ListShardRecoveriesResponseBody {
	return s.Body
}

func (s *ListShardRecoveriesResponse) SetHeaders(v map[string]*string) *ListShardRecoveriesResponse {
	s.Headers = v
	return s
}

func (s *ListShardRecoveriesResponse) SetStatusCode(v int32) *ListShardRecoveriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShardRecoveriesResponse) SetBody(v *ListShardRecoveriesResponseBody) *ListShardRecoveriesResponse {
	s.Body = v
	return s
}

func (s *ListShardRecoveriesResponse) Validate() error {
	return dara.Validate(s)
}
