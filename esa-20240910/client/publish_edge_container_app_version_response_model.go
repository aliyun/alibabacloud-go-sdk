// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishEdgeContainerAppVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishEdgeContainerAppVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishEdgeContainerAppVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishEdgeContainerAppVersionResponseBody) *PublishEdgeContainerAppVersionResponse
	GetBody() *PublishEdgeContainerAppVersionResponseBody
}

type PublishEdgeContainerAppVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishEdgeContainerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishEdgeContainerAppVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishEdgeContainerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishEdgeContainerAppVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishEdgeContainerAppVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishEdgeContainerAppVersionResponse) GetBody() *PublishEdgeContainerAppVersionResponseBody {
	return s.Body
}

func (s *PublishEdgeContainerAppVersionResponse) SetHeaders(v map[string]*string) *PublishEdgeContainerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishEdgeContainerAppVersionResponse) SetStatusCode(v int32) *PublishEdgeContainerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishEdgeContainerAppVersionResponse) SetBody(v *PublishEdgeContainerAppVersionResponseBody) *PublishEdgeContainerAppVersionResponse {
	s.Body = v
	return s
}

func (s *PublishEdgeContainerAppVersionResponse) Validate() error {
	return dara.Validate(s)
}
