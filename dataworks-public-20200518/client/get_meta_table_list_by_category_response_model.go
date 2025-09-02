// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableListByCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableListByCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableListByCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableListByCategoryResponseBody) *GetMetaTableListByCategoryResponse
	GetBody() *GetMetaTableListByCategoryResponseBody
}

type GetMetaTableListByCategoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableListByCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableListByCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableListByCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableListByCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableListByCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableListByCategoryResponse) GetBody() *GetMetaTableListByCategoryResponseBody {
	return s.Body
}

func (s *GetMetaTableListByCategoryResponse) SetHeaders(v map[string]*string) *GetMetaTableListByCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableListByCategoryResponse) SetStatusCode(v int32) *GetMetaTableListByCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableListByCategoryResponse) SetBody(v *GetMetaTableListByCategoryResponseBody) *GetMetaTableListByCategoryResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableListByCategoryResponse) Validate() error {
	return dara.Validate(s)
}
