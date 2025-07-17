// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumUploadFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRumUploadFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRumUploadFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRumUploadFileResponseBody) *DeleteRumUploadFileResponse
	GetBody() *DeleteRumUploadFileResponseBody
}

type DeleteRumUploadFileResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRumUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRumUploadFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumUploadFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteRumUploadFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRumUploadFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRumUploadFileResponse) GetBody() *DeleteRumUploadFileResponseBody {
	return s.Body
}

func (s *DeleteRumUploadFileResponse) SetHeaders(v map[string]*string) *DeleteRumUploadFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteRumUploadFileResponse) SetStatusCode(v int32) *DeleteRumUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRumUploadFileResponse) SetBody(v *DeleteRumUploadFileResponseBody) *DeleteRumUploadFileResponse {
	s.Body = v
	return s
}

func (s *DeleteRumUploadFileResponse) Validate() error {
	return dara.Validate(s)
}
