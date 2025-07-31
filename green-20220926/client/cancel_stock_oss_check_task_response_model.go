// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelStockOssCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelStockOssCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelStockOssCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelStockOssCheckTaskResponseBody) *CancelStockOssCheckTaskResponse
	GetBody() *CancelStockOssCheckTaskResponseBody
}

type CancelStockOssCheckTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelStockOssCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelStockOssCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelStockOssCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelStockOssCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelStockOssCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelStockOssCheckTaskResponse) GetBody() *CancelStockOssCheckTaskResponseBody {
	return s.Body
}

func (s *CancelStockOssCheckTaskResponse) SetHeaders(v map[string]*string) *CancelStockOssCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelStockOssCheckTaskResponse) SetStatusCode(v int32) *CancelStockOssCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelStockOssCheckTaskResponse) SetBody(v *CancelStockOssCheckTaskResponseBody) *CancelStockOssCheckTaskResponse {
	s.Body = v
	return s
}

func (s *CancelStockOssCheckTaskResponse) Validate() error {
	return dara.Validate(s)
}
