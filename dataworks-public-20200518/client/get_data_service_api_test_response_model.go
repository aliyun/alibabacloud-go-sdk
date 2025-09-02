// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiTestResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiTestResponseBody) *GetDataServiceApiTestResponse
	GetBody() *GetDataServiceApiTestResponseBody
}

type GetDataServiceApiTestResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiTestResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiTestResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiTestResponse) GetBody() *GetDataServiceApiTestResponseBody {
	return s.Body
}

func (s *GetDataServiceApiTestResponse) SetHeaders(v map[string]*string) *GetDataServiceApiTestResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiTestResponse) SetStatusCode(v int32) *GetDataServiceApiTestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiTestResponse) SetBody(v *GetDataServiceApiTestResponseBody) *GetDataServiceApiTestResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiTestResponse) Validate() error {
	return dara.Validate(s)
}
