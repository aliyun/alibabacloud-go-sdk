// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesForExpressConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagResourcesForExpressConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagResourcesForExpressConnectResponse
	GetStatusCode() *int32
	SetBody(v *TagResourcesForExpressConnectResponseBody) *TagResourcesForExpressConnectResponse
	GetBody() *TagResourcesForExpressConnectResponseBody
}

type TagResourcesForExpressConnectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesForExpressConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesForExpressConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesForExpressConnectResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesForExpressConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagResourcesForExpressConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagResourcesForExpressConnectResponse) GetBody() *TagResourcesForExpressConnectResponseBody {
	return s.Body
}

func (s *TagResourcesForExpressConnectResponse) SetHeaders(v map[string]*string) *TagResourcesForExpressConnectResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesForExpressConnectResponse) SetStatusCode(v int32) *TagResourcesForExpressConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesForExpressConnectResponse) SetBody(v *TagResourcesForExpressConnectResponseBody) *TagResourcesForExpressConnectResponse {
	s.Body = v
	return s
}

func (s *TagResourcesForExpressConnectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
