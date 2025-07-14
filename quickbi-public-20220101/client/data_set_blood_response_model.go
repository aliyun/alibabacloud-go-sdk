// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSetBloodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DataSetBloodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DataSetBloodResponse
	GetStatusCode() *int32
	SetBody(v *DataSetBloodResponseBody) *DataSetBloodResponse
	GetBody() *DataSetBloodResponseBody
}

type DataSetBloodResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataSetBloodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataSetBloodResponse) String() string {
	return dara.Prettify(s)
}

func (s DataSetBloodResponse) GoString() string {
	return s.String()
}

func (s *DataSetBloodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DataSetBloodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DataSetBloodResponse) GetBody() *DataSetBloodResponseBody {
	return s.Body
}

func (s *DataSetBloodResponse) SetHeaders(v map[string]*string) *DataSetBloodResponse {
	s.Headers = v
	return s
}

func (s *DataSetBloodResponse) SetStatusCode(v int32) *DataSetBloodResponse {
	s.StatusCode = &v
	return s
}

func (s *DataSetBloodResponse) SetBody(v *DataSetBloodResponseBody) *DataSetBloodResponse {
	s.Body = v
	return s
}

func (s *DataSetBloodResponse) Validate() error {
	return dara.Validate(s)
}
