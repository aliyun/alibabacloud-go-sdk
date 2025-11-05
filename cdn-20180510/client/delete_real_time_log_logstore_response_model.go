// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRealTimeLogLogstoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRealTimeLogLogstoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRealTimeLogLogstoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRealTimeLogLogstoreResponseBody) *DeleteRealTimeLogLogstoreResponse
	GetBody() *DeleteRealTimeLogLogstoreResponseBody
}

type DeleteRealTimeLogLogstoreResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRealTimeLogLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRealTimeLogLogstoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRealTimeLogLogstoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteRealTimeLogLogstoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRealTimeLogLogstoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRealTimeLogLogstoreResponse) GetBody() *DeleteRealTimeLogLogstoreResponseBody {
	return s.Body
}

func (s *DeleteRealTimeLogLogstoreResponse) SetHeaders(v map[string]*string) *DeleteRealTimeLogLogstoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteRealTimeLogLogstoreResponse) SetStatusCode(v int32) *DeleteRealTimeLogLogstoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRealTimeLogLogstoreResponse) SetBody(v *DeleteRealTimeLogLogstoreResponseBody) *DeleteRealTimeLogLogstoreResponse {
	s.Body = v
	return s
}

func (s *DeleteRealTimeLogLogstoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
