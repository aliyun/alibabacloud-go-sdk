// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableDetailInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableDetailInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableDetailInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableDetailInfoResponseBody) *GetMetaTableDetailInfoResponse
	GetBody() *GetMetaTableDetailInfoResponseBody
}

type GetMetaTableDetailInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableDetailInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableDetailInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableDetailInfoResponse) GetBody() *GetMetaTableDetailInfoResponseBody {
	return s.Body
}

func (s *GetMetaTableDetailInfoResponse) SetHeaders(v map[string]*string) *GetMetaTableDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableDetailInfoResponse) SetStatusCode(v int32) *GetMetaTableDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableDetailInfoResponse) SetBody(v *GetMetaTableDetailInfoResponseBody) *GetMetaTableDetailInfoResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableDetailInfoResponse) Validate() error {
	return dara.Validate(s)
}
