// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *SuspendHoloWarehouseResponseBody) *SuspendHoloWarehouseResponse
	GetBody() *SuspendHoloWarehouseResponseBody
}

type SuspendHoloWarehouseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *SuspendHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendHoloWarehouseResponse) GetBody() *SuspendHoloWarehouseResponseBody {
	return s.Body
}

func (s *SuspendHoloWarehouseResponse) SetHeaders(v map[string]*string) *SuspendHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *SuspendHoloWarehouseResponse) SetStatusCode(v int32) *SuspendHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendHoloWarehouseResponse) SetBody(v *SuspendHoloWarehouseResponseBody) *SuspendHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *SuspendHoloWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
