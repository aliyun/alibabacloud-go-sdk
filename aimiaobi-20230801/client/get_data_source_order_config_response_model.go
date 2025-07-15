// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceOrderConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataSourceOrderConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataSourceOrderConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetDataSourceOrderConfigResponseBody) *GetDataSourceOrderConfigResponse
	GetBody() *GetDataSourceOrderConfigResponseBody
}

type GetDataSourceOrderConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceOrderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceOrderConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceOrderConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataSourceOrderConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataSourceOrderConfigResponse) GetBody() *GetDataSourceOrderConfigResponseBody {
	return s.Body
}

func (s *GetDataSourceOrderConfigResponse) SetHeaders(v map[string]*string) *GetDataSourceOrderConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceOrderConfigResponse) SetStatusCode(v int32) *GetDataSourceOrderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceOrderConfigResponse) SetBody(v *GetDataSourceOrderConfigResponseBody) *GetDataSourceOrderConfigResponse {
	s.Body = v
	return s
}

func (s *GetDataSourceOrderConfigResponse) Validate() error {
	return dara.Validate(s)
}
