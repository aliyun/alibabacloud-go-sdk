// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAdvanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishAdvanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishAdvanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *PublishAdvanceConfigResponseBody) *PublishAdvanceConfigResponse
	GetBody() *PublishAdvanceConfigResponseBody
}

type PublishAdvanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishAdvanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishAdvanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishAdvanceConfigResponse) GetBody() *PublishAdvanceConfigResponseBody {
	return s.Body
}

func (s *PublishAdvanceConfigResponse) SetHeaders(v map[string]*string) *PublishAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishAdvanceConfigResponse) SetStatusCode(v int32) *PublishAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAdvanceConfigResponse) SetBody(v *PublishAdvanceConfigResponseBody) *PublishAdvanceConfigResponse {
	s.Body = v
	return s
}

func (s *PublishAdvanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
