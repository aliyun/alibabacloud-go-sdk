// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIcebergTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIcebergTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIcebergTableResponse
	GetStatusCode() *int32
	SetBody(v *IcebergTable) *GetIcebergTableResponse
	GetBody() *IcebergTable
}

type GetIcebergTableResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IcebergTable      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIcebergTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIcebergTableResponse) GoString() string {
	return s.String()
}

func (s *GetIcebergTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIcebergTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIcebergTableResponse) GetBody() *IcebergTable {
	return s.Body
}

func (s *GetIcebergTableResponse) SetHeaders(v map[string]*string) *GetIcebergTableResponse {
	s.Headers = v
	return s
}

func (s *GetIcebergTableResponse) SetStatusCode(v int32) *GetIcebergTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIcebergTableResponse) SetBody(v *IcebergTable) *GetIcebergTableResponse {
	s.Body = v
	return s
}

func (s *GetIcebergTableResponse) Validate() error {
	return dara.Validate(s)
}
