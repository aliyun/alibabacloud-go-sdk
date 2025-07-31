// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyDataServiceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyDataServiceAppResponse
	GetStatusCode() *int32
	SetBody(v *ApplyDataServiceAppResponseBody) *ApplyDataServiceAppResponse
	GetBody() *ApplyDataServiceAppResponseBody
}

type ApplyDataServiceAppResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyDataServiceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyDataServiceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceAppResponse) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyDataServiceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyDataServiceAppResponse) GetBody() *ApplyDataServiceAppResponseBody {
	return s.Body
}

func (s *ApplyDataServiceAppResponse) SetHeaders(v map[string]*string) *ApplyDataServiceAppResponse {
	s.Headers = v
	return s
}

func (s *ApplyDataServiceAppResponse) SetStatusCode(v int32) *ApplyDataServiceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyDataServiceAppResponse) SetBody(v *ApplyDataServiceAppResponseBody) *ApplyDataServiceAppResponse {
	s.Body = v
	return s
}

func (s *ApplyDataServiceAppResponse) Validate() error {
	return dara.Validate(s)
}
