// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStockOssCheckTasksListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStockOssCheckTasksListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStockOssCheckTasksListResponse
	GetStatusCode() *int32
	SetBody(v *GetStockOssCheckTasksListResponseBody) *GetStockOssCheckTasksListResponse
	GetBody() *GetStockOssCheckTasksListResponseBody
}

type GetStockOssCheckTasksListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStockOssCheckTasksListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStockOssCheckTasksListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListResponse) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStockOssCheckTasksListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStockOssCheckTasksListResponse) GetBody() *GetStockOssCheckTasksListResponseBody {
	return s.Body
}

func (s *GetStockOssCheckTasksListResponse) SetHeaders(v map[string]*string) *GetStockOssCheckTasksListResponse {
	s.Headers = v
	return s
}

func (s *GetStockOssCheckTasksListResponse) SetStatusCode(v int32) *GetStockOssCheckTasksListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStockOssCheckTasksListResponse) SetBody(v *GetStockOssCheckTasksListResponseBody) *GetStockOssCheckTasksListResponse {
	s.Body = v
	return s
}

func (s *GetStockOssCheckTasksListResponse) Validate() error {
	return dara.Validate(s)
}
