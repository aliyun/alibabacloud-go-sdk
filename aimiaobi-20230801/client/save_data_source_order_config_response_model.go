// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDataSourceOrderConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveDataSourceOrderConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveDataSourceOrderConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveDataSourceOrderConfigResponseBody) *SaveDataSourceOrderConfigResponse
	GetBody() *SaveDataSourceOrderConfigResponseBody
}

type SaveDataSourceOrderConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveDataSourceOrderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveDataSourceOrderConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveDataSourceOrderConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveDataSourceOrderConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveDataSourceOrderConfigResponse) GetBody() *SaveDataSourceOrderConfigResponseBody {
	return s.Body
}

func (s *SaveDataSourceOrderConfigResponse) SetHeaders(v map[string]*string) *SaveDataSourceOrderConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) SetStatusCode(v int32) *SaveDataSourceOrderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) SetBody(v *SaveDataSourceOrderConfigResponseBody) *SaveDataSourceOrderConfigResponse {
	s.Body = v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) Validate() error {
	return dara.Validate(s)
}
