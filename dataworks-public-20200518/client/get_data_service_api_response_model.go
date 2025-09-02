// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiResponseBody) *GetDataServiceApiResponse
	GetBody() *GetDataServiceApiResponseBody
}

type GetDataServiceApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiResponse) GetBody() *GetDataServiceApiResponseBody {
	return s.Body
}

func (s *GetDataServiceApiResponse) SetHeaders(v map[string]*string) *GetDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiResponse) SetStatusCode(v int32) *GetDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiResponse) SetBody(v *GetDataServiceApiResponseBody) *GetDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
