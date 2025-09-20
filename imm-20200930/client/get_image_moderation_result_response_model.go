// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetImageModerationResultResponseBody) *GetImageModerationResultResponse
	GetBody() *GetImageModerationResultResponseBody
}

type GetImageModerationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageModerationResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageModerationResultResponse) GetBody() *GetImageModerationResultResponseBody {
	return s.Body
}

func (s *GetImageModerationResultResponse) SetHeaders(v map[string]*string) *GetImageModerationResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageModerationResultResponse) SetStatusCode(v int32) *GetImageModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageModerationResultResponse) SetBody(v *GetImageModerationResultResponseBody) *GetImageModerationResultResponse {
	s.Body = v
	return s
}

func (s *GetImageModerationResultResponse) Validate() error {
	return dara.Validate(s)
}
