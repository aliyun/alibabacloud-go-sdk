// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJiangSuTelecomDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJiangSuTelecomDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJiangSuTelecomDataResponse
	GetStatusCode() *int32
	SetBody(v *GetJiangSuTelecomDataResponseBody) *GetJiangSuTelecomDataResponse
	GetBody() *GetJiangSuTelecomDataResponseBody
}

type GetJiangSuTelecomDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJiangSuTelecomDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJiangSuTelecomDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJiangSuTelecomDataResponse) GoString() string {
	return s.String()
}

func (s *GetJiangSuTelecomDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJiangSuTelecomDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJiangSuTelecomDataResponse) GetBody() *GetJiangSuTelecomDataResponseBody {
	return s.Body
}

func (s *GetJiangSuTelecomDataResponse) SetHeaders(v map[string]*string) *GetJiangSuTelecomDataResponse {
	s.Headers = v
	return s
}

func (s *GetJiangSuTelecomDataResponse) SetStatusCode(v int32) *GetJiangSuTelecomDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJiangSuTelecomDataResponse) SetBody(v *GetJiangSuTelecomDataResponseBody) *GetJiangSuTelecomDataResponse {
	s.Body = v
	return s
}

func (s *GetJiangSuTelecomDataResponse) Validate() error {
	return dara.Validate(s)
}
