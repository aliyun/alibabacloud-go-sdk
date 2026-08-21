// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRecordWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateZoneRecordWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateZoneRecordWeightResponse
	GetStatusCode() *int32
	SetBody(v *UpdateZoneRecordWeightResponseBody) *UpdateZoneRecordWeightResponse
	GetBody() *UpdateZoneRecordWeightResponseBody
}

type UpdateZoneRecordWeightResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZoneRecordWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZoneRecordWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRecordWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateZoneRecordWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateZoneRecordWeightResponse) GetBody() *UpdateZoneRecordWeightResponseBody {
	return s.Body
}

func (s *UpdateZoneRecordWeightResponse) SetHeaders(v map[string]*string) *UpdateZoneRecordWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRecordWeightResponse) SetStatusCode(v int32) *UpdateZoneRecordWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZoneRecordWeightResponse) SetBody(v *UpdateZoneRecordWeightResponseBody) *UpdateZoneRecordWeightResponse {
	s.Body = v
	return s
}

func (s *UpdateZoneRecordWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
