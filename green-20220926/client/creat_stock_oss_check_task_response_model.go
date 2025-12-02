// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatStockOssCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatStockOssCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatStockOssCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatStockOssCheckTaskResponseBody) *CreatStockOssCheckTaskResponse
	GetBody() *CreatStockOssCheckTaskResponseBody
}

type CreatStockOssCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatStockOssCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatStockOssCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatStockOssCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatStockOssCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatStockOssCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatStockOssCheckTaskResponse) GetBody() *CreatStockOssCheckTaskResponseBody {
	return s.Body
}

func (s *CreatStockOssCheckTaskResponse) SetHeaders(v map[string]*string) *CreatStockOssCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatStockOssCheckTaskResponse) SetStatusCode(v int32) *CreatStockOssCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatStockOssCheckTaskResponse) SetBody(v *CreatStockOssCheckTaskResponseBody) *CreatStockOssCheckTaskResponse {
	s.Body = v
	return s
}

func (s *CreatStockOssCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
