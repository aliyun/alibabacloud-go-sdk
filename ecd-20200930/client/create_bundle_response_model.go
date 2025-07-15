// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBundleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBundleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBundleResponse
	GetStatusCode() *int32
	SetBody(v *CreateBundleResponseBody) *CreateBundleResponse
	GetBody() *CreateBundleResponseBody
}

type CreateBundleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBundleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBundleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateBundleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBundleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBundleResponse) GetBody() *CreateBundleResponseBody {
	return s.Body
}

func (s *CreateBundleResponse) SetHeaders(v map[string]*string) *CreateBundleResponse {
	s.Headers = v
	return s
}

func (s *CreateBundleResponse) SetStatusCode(v int32) *CreateBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBundleResponse) SetBody(v *CreateBundleResponseBody) *CreateBundleResponse {
	s.Body = v
	return s
}

func (s *CreateBundleResponse) Validate() error {
	return dara.Validate(s)
}
