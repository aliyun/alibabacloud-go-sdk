// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCorrectOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCorrectOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCorrectOrderDetailResponseBody) *GetDataCorrectOrderDetailResponse
	GetBody() *GetDataCorrectOrderDetailResponseBody
}

type GetDataCorrectOrderDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCorrectOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCorrectOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCorrectOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCorrectOrderDetailResponse) GetBody() *GetDataCorrectOrderDetailResponseBody {
	return s.Body
}

func (s *GetDataCorrectOrderDetailResponse) SetHeaders(v map[string]*string) *GetDataCorrectOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectOrderDetailResponse) SetStatusCode(v int32) *GetDataCorrectOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponse) SetBody(v *GetDataCorrectOrderDetailResponseBody) *GetDataCorrectOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetDataCorrectOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
