// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleWithAdjustmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleWithAdjustmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleWithAdjustmentResponse
	GetStatusCode() *int32
	SetBody(v *ScaleWithAdjustmentResponseBody) *ScaleWithAdjustmentResponse
	GetBody() *ScaleWithAdjustmentResponseBody
}

type ScaleWithAdjustmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleWithAdjustmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleWithAdjustmentResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentResponse) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleWithAdjustmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleWithAdjustmentResponse) GetBody() *ScaleWithAdjustmentResponseBody {
	return s.Body
}

func (s *ScaleWithAdjustmentResponse) SetHeaders(v map[string]*string) *ScaleWithAdjustmentResponse {
	s.Headers = v
	return s
}

func (s *ScaleWithAdjustmentResponse) SetStatusCode(v int32) *ScaleWithAdjustmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleWithAdjustmentResponse) SetBody(v *ScaleWithAdjustmentResponseBody) *ScaleWithAdjustmentResponse {
	s.Body = v
	return s
}

func (s *ScaleWithAdjustmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
