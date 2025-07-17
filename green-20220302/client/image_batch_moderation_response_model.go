// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBatchModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageBatchModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageBatchModerationResponse
	GetStatusCode() *int32
	SetBody(v *ImageBatchModerationResponseBody) *ImageBatchModerationResponse
	GetBody() *ImageBatchModerationResponseBody
}

type ImageBatchModerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageBatchModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageBatchModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageBatchModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageBatchModerationResponse) GetBody() *ImageBatchModerationResponseBody {
	return s.Body
}

func (s *ImageBatchModerationResponse) SetHeaders(v map[string]*string) *ImageBatchModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageBatchModerationResponse) SetStatusCode(v int32) *ImageBatchModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageBatchModerationResponse) SetBody(v *ImageBatchModerationResponseBody) *ImageBatchModerationResponse {
	s.Body = v
	return s
}

func (s *ImageBatchModerationResponse) Validate() error {
	return dara.Validate(s)
}
