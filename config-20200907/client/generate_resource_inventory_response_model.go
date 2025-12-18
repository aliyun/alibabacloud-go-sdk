// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourceInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateResourceInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateResourceInventoryResponse
	GetStatusCode() *int32
	SetBody(v *GenerateResourceInventoryResponseBody) *GenerateResourceInventoryResponse
	GetBody() *GenerateResourceInventoryResponseBody
}

type GenerateResourceInventoryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateResourceInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateResourceInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourceInventoryResponse) GoString() string {
	return s.String()
}

func (s *GenerateResourceInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateResourceInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateResourceInventoryResponse) GetBody() *GenerateResourceInventoryResponseBody {
	return s.Body
}

func (s *GenerateResourceInventoryResponse) SetHeaders(v map[string]*string) *GenerateResourceInventoryResponse {
	s.Headers = v
	return s
}

func (s *GenerateResourceInventoryResponse) SetStatusCode(v int32) *GenerateResourceInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateResourceInventoryResponse) SetBody(v *GenerateResourceInventoryResponseBody) *GenerateResourceInventoryResponse {
	s.Body = v
	return s
}

func (s *GenerateResourceInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
