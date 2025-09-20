// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConvertTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaConvertTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaConvertTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaConvertTaskResponseBody) *CreateMediaConvertTaskResponse
	GetBody() *CreateMediaConvertTaskResponseBody
}

type CreateMediaConvertTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaConvertTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaConvertTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaConvertTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaConvertTaskResponse) GetBody() *CreateMediaConvertTaskResponseBody {
	return s.Body
}

func (s *CreateMediaConvertTaskResponse) SetHeaders(v map[string]*string) *CreateMediaConvertTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaConvertTaskResponse) SetStatusCode(v int32) *CreateMediaConvertTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaConvertTaskResponse) SetBody(v *CreateMediaConvertTaskResponseBody) *CreateMediaConvertTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMediaConvertTaskResponse) Validate() error {
	return dara.Validate(s)
}
