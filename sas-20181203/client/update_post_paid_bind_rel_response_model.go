// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostPaidBindRelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePostPaidBindRelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePostPaidBindRelResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePostPaidBindRelResponseBody) *UpdatePostPaidBindRelResponse
	GetBody() *UpdatePostPaidBindRelResponseBody
}

type UpdatePostPaidBindRelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePostPaidBindRelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePostPaidBindRelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelResponse) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePostPaidBindRelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePostPaidBindRelResponse) GetBody() *UpdatePostPaidBindRelResponseBody {
	return s.Body
}

func (s *UpdatePostPaidBindRelResponse) SetHeaders(v map[string]*string) *UpdatePostPaidBindRelResponse {
	s.Headers = v
	return s
}

func (s *UpdatePostPaidBindRelResponse) SetStatusCode(v int32) *UpdatePostPaidBindRelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePostPaidBindRelResponse) SetBody(v *UpdatePostPaidBindRelResponseBody) *UpdatePostPaidBindRelResponse {
	s.Body = v
	return s
}

func (s *UpdatePostPaidBindRelResponse) Validate() error {
	return dara.Validate(s)
}
