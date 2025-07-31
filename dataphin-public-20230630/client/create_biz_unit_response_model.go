// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBizUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBizUnitResponse
	GetStatusCode() *int32
	SetBody(v *CreateBizUnitResponseBody) *CreateBizUnitResponse
	GetBody() *CreateBizUnitResponseBody
}

type CreateBizUnitResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBizUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBizUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateBizUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBizUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBizUnitResponse) GetBody() *CreateBizUnitResponseBody {
	return s.Body
}

func (s *CreateBizUnitResponse) SetHeaders(v map[string]*string) *CreateBizUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateBizUnitResponse) SetStatusCode(v int32) *CreateBizUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBizUnitResponse) SetBody(v *CreateBizUnitResponseBody) *CreateBizUnitResponse {
	s.Body = v
	return s
}

func (s *CreateBizUnitResponse) Validate() error {
	return dara.Validate(s)
}
