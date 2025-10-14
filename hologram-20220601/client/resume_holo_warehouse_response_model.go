// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *ResumeHoloWarehouseResponseBody) *ResumeHoloWarehouseResponse
	GetBody() *ResumeHoloWarehouseResponseBody
}

type ResumeHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *ResumeHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeHoloWarehouseResponse) GetBody() *ResumeHoloWarehouseResponseBody {
	return s.Body
}

func (s *ResumeHoloWarehouseResponse) SetHeaders(v map[string]*string) *ResumeHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *ResumeHoloWarehouseResponse) SetStatusCode(v int32) *ResumeHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeHoloWarehouseResponse) SetBody(v *ResumeHoloWarehouseResponseBody) *ResumeHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *ResumeHoloWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
