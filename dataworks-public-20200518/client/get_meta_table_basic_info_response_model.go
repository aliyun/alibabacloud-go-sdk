// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableBasicInfoResponseBody) *GetMetaTableBasicInfoResponse
	GetBody() *GetMetaTableBasicInfoResponseBody
}

type GetMetaTableBasicInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableBasicInfoResponse) GetBody() *GetMetaTableBasicInfoResponseBody {
	return s.Body
}

func (s *GetMetaTableBasicInfoResponse) SetHeaders(v map[string]*string) *GetMetaTableBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableBasicInfoResponse) SetStatusCode(v int32) *GetMetaTableBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableBasicInfoResponse) SetBody(v *GetMetaTableBasicInfoResponseBody) *GetMetaTableBasicInfoResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableBasicInfoResponse) Validate() error {
	return dara.Validate(s)
}
