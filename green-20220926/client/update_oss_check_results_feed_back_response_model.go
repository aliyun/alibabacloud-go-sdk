// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsFeedBackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOssCheckResultsFeedBackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOssCheckResultsFeedBackResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOssCheckResultsFeedBackResponseBody) *UpdateOssCheckResultsFeedBackResponse
	GetBody() *UpdateOssCheckResultsFeedBackResponseBody
}

type UpdateOssCheckResultsFeedBackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssCheckResultsFeedBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssCheckResultsFeedBackResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFeedBackResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFeedBackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOssCheckResultsFeedBackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOssCheckResultsFeedBackResponse) GetBody() *UpdateOssCheckResultsFeedBackResponseBody {
	return s.Body
}

func (s *UpdateOssCheckResultsFeedBackResponse) SetHeaders(v map[string]*string) *UpdateOssCheckResultsFeedBackResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssCheckResultsFeedBackResponse) SetStatusCode(v int32) *UpdateOssCheckResultsFeedBackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackResponse) SetBody(v *UpdateOssCheckResultsFeedBackResponseBody) *UpdateOssCheckResultsFeedBackResponse {
	s.Body = v
	return s
}

func (s *UpdateOssCheckResultsFeedBackResponse) Validate() error {
	return dara.Validate(s)
}
