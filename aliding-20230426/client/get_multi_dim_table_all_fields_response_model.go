// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiDimTableAllFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiDimTableAllFieldsResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiDimTableAllFieldsResponseBody) *GetMultiDimTableAllFieldsResponse
	GetBody() *GetMultiDimTableAllFieldsResponseBody
}

type GetMultiDimTableAllFieldsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiDimTableAllFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiDimTableAllFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsResponse) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiDimTableAllFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiDimTableAllFieldsResponse) GetBody() *GetMultiDimTableAllFieldsResponseBody {
	return s.Body
}

func (s *GetMultiDimTableAllFieldsResponse) SetHeaders(v map[string]*string) *GetMultiDimTableAllFieldsResponse {
	s.Headers = v
	return s
}

func (s *GetMultiDimTableAllFieldsResponse) SetStatusCode(v int32) *GetMultiDimTableAllFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponse) SetBody(v *GetMultiDimTableAllFieldsResponseBody) *GetMultiDimTableAllFieldsResponse {
	s.Body = v
	return s
}

func (s *GetMultiDimTableAllFieldsResponse) Validate() error {
	return dara.Validate(s)
}
