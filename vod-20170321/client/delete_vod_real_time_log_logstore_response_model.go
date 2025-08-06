// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodRealTimeLogLogstoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodRealTimeLogLogstoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodRealTimeLogLogstoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodRealTimeLogLogstoreResponseBody) *DeleteVodRealTimeLogLogstoreResponse
	GetBody() *DeleteVodRealTimeLogLogstoreResponseBody
}

type DeleteVodRealTimeLogLogstoreResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodRealTimeLogLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodRealTimeLogLogstoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodRealTimeLogLogstoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodRealTimeLogLogstoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodRealTimeLogLogstoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodRealTimeLogLogstoreResponse) GetBody() *DeleteVodRealTimeLogLogstoreResponseBody {
	return s.Body
}

func (s *DeleteVodRealTimeLogLogstoreResponse) SetHeaders(v map[string]*string) *DeleteVodRealTimeLogLogstoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreResponse) SetStatusCode(v int32) *DeleteVodRealTimeLogLogstoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreResponse) SetBody(v *DeleteVodRealTimeLogLogstoreResponseBody) *DeleteVodRealTimeLogLogstoreResponse {
	s.Body = v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreResponse) Validate() error {
	return dara.Validate(s)
}
