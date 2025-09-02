// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMeasureDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMeasureDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMeasureDataResponse
	GetStatusCode() *int32
	SetBody(v *ListMeasureDataResponseBody) *ListMeasureDataResponse
	GetBody() *ListMeasureDataResponseBody
}

type ListMeasureDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMeasureDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *ListMeasureDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMeasureDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMeasureDataResponse) GetBody() *ListMeasureDataResponseBody {
	return s.Body
}

func (s *ListMeasureDataResponse) SetHeaders(v map[string]*string) *ListMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *ListMeasureDataResponse) SetStatusCode(v int32) *ListMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMeasureDataResponse) SetBody(v *ListMeasureDataResponseBody) *ListMeasureDataResponse {
	s.Body = v
	return s
}

func (s *ListMeasureDataResponse) Validate() error {
	return dara.Validate(s)
}
