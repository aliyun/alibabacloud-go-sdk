// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesForExpressConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagResourcesForExpressConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagResourcesForExpressConnectResponse
	GetStatusCode() *int32
	SetBody(v *UntagResourcesForExpressConnectResponseBody) *UntagResourcesForExpressConnectResponse
	GetBody() *UntagResourcesForExpressConnectResponseBody
}

type UntagResourcesForExpressConnectResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesForExpressConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesForExpressConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesForExpressConnectResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesForExpressConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagResourcesForExpressConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagResourcesForExpressConnectResponse) GetBody() *UntagResourcesForExpressConnectResponseBody {
	return s.Body
}

func (s *UntagResourcesForExpressConnectResponse) SetHeaders(v map[string]*string) *UntagResourcesForExpressConnectResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesForExpressConnectResponse) SetStatusCode(v int32) *UntagResourcesForExpressConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesForExpressConnectResponse) SetBody(v *UntagResourcesForExpressConnectResponseBody) *UntagResourcesForExpressConnectResponse {
	s.Body = v
	return s
}

func (s *UntagResourcesForExpressConnectResponse) Validate() error {
	return dara.Validate(s)
}
