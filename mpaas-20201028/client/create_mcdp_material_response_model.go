// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcdpMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcdpMaterialResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcdpMaterialResponseBody) *CreateMcdpMaterialResponse
	GetBody() *CreateMcdpMaterialResponseBody
}

type CreateMcdpMaterialResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcdpMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcdpMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpMaterialResponse) GoString() string {
	return s.String()
}

func (s *CreateMcdpMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcdpMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcdpMaterialResponse) GetBody() *CreateMcdpMaterialResponseBody {
	return s.Body
}

func (s *CreateMcdpMaterialResponse) SetHeaders(v map[string]*string) *CreateMcdpMaterialResponse {
	s.Headers = v
	return s
}

func (s *CreateMcdpMaterialResponse) SetStatusCode(v int32) *CreateMcdpMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcdpMaterialResponse) SetBody(v *CreateMcdpMaterialResponseBody) *CreateMcdpMaterialResponse {
	s.Body = v
	return s
}

func (s *CreateMcdpMaterialResponse) Validate() error {
	return dara.Validate(s)
}
