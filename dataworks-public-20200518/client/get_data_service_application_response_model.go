// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApplicationResponseBody) *GetDataServiceApplicationResponse
	GetBody() *GetDataServiceApplicationResponseBody
}

type GetDataServiceApplicationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApplicationResponse) GetBody() *GetDataServiceApplicationResponseBody {
	return s.Body
}

func (s *GetDataServiceApplicationResponse) SetHeaders(v map[string]*string) *GetDataServiceApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApplicationResponse) SetStatusCode(v int32) *GetDataServiceApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApplicationResponse) SetBody(v *GetDataServiceApplicationResponseBody) *GetDataServiceApplicationResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApplicationResponse) Validate() error {
	return dara.Validate(s)
}
