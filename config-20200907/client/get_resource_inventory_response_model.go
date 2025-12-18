// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceInventoryResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceInventoryResponseBody) *GetResourceInventoryResponse
	GetBody() *GetResourceInventoryResponseBody
}

type GetResourceInventoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceInventoryResponse) GoString() string {
	return s.String()
}

func (s *GetResourceInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceInventoryResponse) GetBody() *GetResourceInventoryResponseBody {
	return s.Body
}

func (s *GetResourceInventoryResponse) SetHeaders(v map[string]*string) *GetResourceInventoryResponse {
	s.Headers = v
	return s
}

func (s *GetResourceInventoryResponse) SetStatusCode(v int32) *GetResourceInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceInventoryResponse) SetBody(v *GetResourceInventoryResponseBody) *GetResourceInventoryResponse {
	s.Body = v
	return s
}

func (s *GetResourceInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
