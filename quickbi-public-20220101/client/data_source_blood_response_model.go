// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSourceBloodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DataSourceBloodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DataSourceBloodResponse
	GetStatusCode() *int32
	SetBody(v *DataSourceBloodResponseBody) *DataSourceBloodResponse
	GetBody() *DataSourceBloodResponseBody
}

type DataSourceBloodResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataSourceBloodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataSourceBloodResponse) String() string {
	return dara.Prettify(s)
}

func (s DataSourceBloodResponse) GoString() string {
	return s.String()
}

func (s *DataSourceBloodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DataSourceBloodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DataSourceBloodResponse) GetBody() *DataSourceBloodResponseBody {
	return s.Body
}

func (s *DataSourceBloodResponse) SetHeaders(v map[string]*string) *DataSourceBloodResponse {
	s.Headers = v
	return s
}

func (s *DataSourceBloodResponse) SetStatusCode(v int32) *DataSourceBloodResponse {
	s.StatusCode = &v
	return s
}

func (s *DataSourceBloodResponse) SetBody(v *DataSourceBloodResponseBody) *DataSourceBloodResponse {
	s.Body = v
	return s
}

func (s *DataSourceBloodResponse) Validate() error {
	return dara.Validate(s)
}
