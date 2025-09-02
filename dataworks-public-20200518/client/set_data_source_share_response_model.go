// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataSourceShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDataSourceShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDataSourceShareResponse
	GetStatusCode() *int32
	SetBody(v *SetDataSourceShareResponseBody) *SetDataSourceShareResponse
	GetBody() *SetDataSourceShareResponseBody
}

type SetDataSourceShareResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataSourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDataSourceShareResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDataSourceShareResponse) GoString() string {
	return s.String()
}

func (s *SetDataSourceShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDataSourceShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDataSourceShareResponse) GetBody() *SetDataSourceShareResponseBody {
	return s.Body
}

func (s *SetDataSourceShareResponse) SetHeaders(v map[string]*string) *SetDataSourceShareResponse {
	s.Headers = v
	return s
}

func (s *SetDataSourceShareResponse) SetStatusCode(v int32) *SetDataSourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataSourceShareResponse) SetBody(v *SetDataSourceShareResponseBody) *SetDataSourceShareResponse {
	s.Body = v
	return s
}

func (s *SetDataSourceShareResponse) Validate() error {
	return dara.Validate(s)
}
