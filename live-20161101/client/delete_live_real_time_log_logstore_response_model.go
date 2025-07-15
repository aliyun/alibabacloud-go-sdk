// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRealTimeLogLogstoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveRealTimeLogLogstoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveRealTimeLogLogstoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveRealTimeLogLogstoreResponseBody) *DeleteLiveRealTimeLogLogstoreResponse
	GetBody() *DeleteLiveRealTimeLogLogstoreResponseBody
}

type DeleteLiveRealTimeLogLogstoreResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveRealTimeLogLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveRealTimeLogLogstoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRealTimeLogLogstoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) GetBody() *DeleteLiveRealTimeLogLogstoreResponseBody {
	return s.Body
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) SetHeaders(v map[string]*string) *DeleteLiveRealTimeLogLogstoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) SetStatusCode(v int32) *DeleteLiveRealTimeLogLogstoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) SetBody(v *DeleteLiveRealTimeLogLogstoreResponseBody) *DeleteLiveRealTimeLogLogstoreResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) Validate() error {
	return dara.Validate(s)
}
