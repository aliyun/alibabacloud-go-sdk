// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishFlowVersionResponseBody) *PublishFlowVersionResponse
	GetBody() *PublishFlowVersionResponseBody
}

type PublishFlowVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishFlowVersionResponse) GetBody() *PublishFlowVersionResponseBody {
	return s.Body
}

func (s *PublishFlowVersionResponse) SetHeaders(v map[string]*string) *PublishFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishFlowVersionResponse) SetStatusCode(v int32) *PublishFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFlowVersionResponse) SetBody(v *PublishFlowVersionResponseBody) *PublishFlowVersionResponse {
	s.Body = v
	return s
}

func (s *PublishFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
