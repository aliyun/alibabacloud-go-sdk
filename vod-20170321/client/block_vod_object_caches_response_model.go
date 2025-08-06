// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockVodObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BlockVodObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BlockVodObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *BlockVodObjectCachesResponseBody) *BlockVodObjectCachesResponse
	GetBody() *BlockVodObjectCachesResponseBody
}

type BlockVodObjectCachesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BlockVodObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BlockVodObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s BlockVodObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *BlockVodObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BlockVodObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BlockVodObjectCachesResponse) GetBody() *BlockVodObjectCachesResponseBody {
	return s.Body
}

func (s *BlockVodObjectCachesResponse) SetHeaders(v map[string]*string) *BlockVodObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *BlockVodObjectCachesResponse) SetStatusCode(v int32) *BlockVodObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *BlockVodObjectCachesResponse) SetBody(v *BlockVodObjectCachesResponseBody) *BlockVodObjectCachesResponse {
	s.Body = v
	return s
}

func (s *BlockVodObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
