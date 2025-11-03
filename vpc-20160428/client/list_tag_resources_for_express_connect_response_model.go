// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesForExpressConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagResourcesForExpressConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagResourcesForExpressConnectResponse
	GetStatusCode() *int32
	SetBody(v *ListTagResourcesForExpressConnectResponseBody) *ListTagResourcesForExpressConnectResponse
	GetBody() *ListTagResourcesForExpressConnectResponseBody
}

type ListTagResourcesForExpressConnectResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesForExpressConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesForExpressConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForExpressConnectResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForExpressConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagResourcesForExpressConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagResourcesForExpressConnectResponse) GetBody() *ListTagResourcesForExpressConnectResponseBody {
	return s.Body
}

func (s *ListTagResourcesForExpressConnectResponse) SetHeaders(v map[string]*string) *ListTagResourcesForExpressConnectResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesForExpressConnectResponse) SetStatusCode(v int32) *ListTagResourcesForExpressConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesForExpressConnectResponse) SetBody(v *ListTagResourcesForExpressConnectResponseBody) *ListTagResourcesForExpressConnectResponse {
	s.Body = v
	return s
}

func (s *ListTagResourcesForExpressConnectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
