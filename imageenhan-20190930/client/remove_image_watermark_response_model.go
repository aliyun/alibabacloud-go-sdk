// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveImageWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveImageWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *RemoveImageWatermarkResponseBody) *RemoveImageWatermarkResponse
	GetBody() *RemoveImageWatermarkResponseBody
}

type RemoveImageWatermarkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageWatermarkResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveImageWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveImageWatermarkResponse) GetBody() *RemoveImageWatermarkResponseBody {
	return s.Body
}

func (s *RemoveImageWatermarkResponse) SetHeaders(v map[string]*string) *RemoveImageWatermarkResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageWatermarkResponse) SetStatusCode(v int32) *RemoveImageWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageWatermarkResponse) SetBody(v *RemoveImageWatermarkResponseBody) *RemoveImageWatermarkResponse {
	s.Body = v
	return s
}

func (s *RemoveImageWatermarkResponse) Validate() error {
	return dara.Validate(s)
}
