// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConvertJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaConvertJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaConvertJobResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaConvertJobResponseBody) *GetMediaConvertJobResponse
	GetBody() *GetMediaConvertJobResponseBody
}

type GetMediaConvertJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaConvertJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaConvertJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConvertJobResponse) GoString() string {
	return s.String()
}

func (s *GetMediaConvertJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaConvertJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaConvertJobResponse) GetBody() *GetMediaConvertJobResponseBody {
	return s.Body
}

func (s *GetMediaConvertJobResponse) SetHeaders(v map[string]*string) *GetMediaConvertJobResponse {
	s.Headers = v
	return s
}

func (s *GetMediaConvertJobResponse) SetStatusCode(v int32) *GetMediaConvertJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaConvertJobResponse) SetBody(v *GetMediaConvertJobResponseBody) *GetMediaConvertJobResponse {
	s.Body = v
	return s
}

func (s *GetMediaConvertJobResponse) Validate() error {
	return dara.Validate(s)
}
