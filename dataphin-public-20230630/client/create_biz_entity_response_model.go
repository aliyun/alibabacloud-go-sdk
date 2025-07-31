// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBizEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBizEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateBizEntityResponseBody) *CreateBizEntityResponse
	GetBody() *CreateBizEntityResponseBody
}

type CreateBizEntityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBizEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBizEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateBizEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBizEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBizEntityResponse) GetBody() *CreateBizEntityResponseBody {
	return s.Body
}

func (s *CreateBizEntityResponse) SetHeaders(v map[string]*string) *CreateBizEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateBizEntityResponse) SetStatusCode(v int32) *CreateBizEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBizEntityResponse) SetBody(v *CreateBizEntityResponseBody) *CreateBizEntityResponse {
	s.Body = v
	return s
}

func (s *CreateBizEntityResponse) Validate() error {
	return dara.Validate(s)
}
