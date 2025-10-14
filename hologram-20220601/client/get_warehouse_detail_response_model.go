// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarehouseDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWarehouseDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWarehouseDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetWarehouseDetailResponseBody) *GetWarehouseDetailResponse
	GetBody() *GetWarehouseDetailResponseBody
}

type GetWarehouseDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWarehouseDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWarehouseDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponse) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWarehouseDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWarehouseDetailResponse) GetBody() *GetWarehouseDetailResponseBody {
	return s.Body
}

func (s *GetWarehouseDetailResponse) SetHeaders(v map[string]*string) *GetWarehouseDetailResponse {
	s.Headers = v
	return s
}

func (s *GetWarehouseDetailResponse) SetStatusCode(v int32) *GetWarehouseDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWarehouseDetailResponse) SetBody(v *GetWarehouseDetailResponseBody) *GetWarehouseDetailResponse {
	s.Body = v
	return s
}

func (s *GetWarehouseDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
