// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishCronResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublishCronResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublishCronResponse
	GetStatusCode() *int32
	SetBody(v *GetPublishCronResponseBody) *GetPublishCronResponse
	GetBody() *GetPublishCronResponseBody
}

type GetPublishCronResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublishCronResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublishCronResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublishCronResponse) GoString() string {
	return s.String()
}

func (s *GetPublishCronResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublishCronResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublishCronResponse) GetBody() *GetPublishCronResponseBody {
	return s.Body
}

func (s *GetPublishCronResponse) SetHeaders(v map[string]*string) *GetPublishCronResponse {
	s.Headers = v
	return s
}

func (s *GetPublishCronResponse) SetStatusCode(v int32) *GetPublishCronResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublishCronResponse) SetBody(v *GetPublishCronResponseBody) *GetPublishCronResponse {
	s.Body = v
	return s
}

func (s *GetPublishCronResponse) Validate() error {
	return dara.Validate(s)
}
