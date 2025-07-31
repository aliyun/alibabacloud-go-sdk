// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdHocFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdHocFileResponse
	GetStatusCode() *int32
	SetBody(v *GetAdHocFileResponseBody) *GetAdHocFileResponse
	GetBody() *GetAdHocFileResponseBody
}

type GetAdHocFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdHocFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *GetAdHocFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdHocFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdHocFileResponse) GetBody() *GetAdHocFileResponseBody {
	return s.Body
}

func (s *GetAdHocFileResponse) SetHeaders(v map[string]*string) *GetAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *GetAdHocFileResponse) SetStatusCode(v int32) *GetAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdHocFileResponse) SetBody(v *GetAdHocFileResponseBody) *GetAdHocFileResponse {
	s.Body = v
	return s
}

func (s *GetAdHocFileResponse) Validate() error {
	return dara.Validate(s)
}
