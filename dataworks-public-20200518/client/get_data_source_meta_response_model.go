// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataSourceMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataSourceMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetDataSourceMetaResponseBody) *GetDataSourceMetaResponse
	GetBody() *GetDataSourceMetaResponseBody
}

type GetDataSourceMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceMetaResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataSourceMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataSourceMetaResponse) GetBody() *GetDataSourceMetaResponseBody {
	return s.Body
}

func (s *GetDataSourceMetaResponse) SetHeaders(v map[string]*string) *GetDataSourceMetaResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceMetaResponse) SetStatusCode(v int32) *GetDataSourceMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceMetaResponse) SetBody(v *GetDataSourceMetaResponseBody) *GetDataSourceMetaResponse {
	s.Body = v
	return s
}

func (s *GetDataSourceMetaResponse) Validate() error {
	return dara.Validate(s)
}
