// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityCheckDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQualityCheckDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQualityCheckDataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQualityCheckDataResponseBody) *UpdateQualityCheckDataResponse
	GetBody() *UpdateQualityCheckDataResponseBody
}

type UpdateQualityCheckDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityCheckDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQualityCheckDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQualityCheckDataResponse) GetBody() *UpdateQualityCheckDataResponseBody {
	return s.Body
}

func (s *UpdateQualityCheckDataResponse) SetHeaders(v map[string]*string) *UpdateQualityCheckDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityCheckDataResponse) SetStatusCode(v int32) *UpdateQualityCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityCheckDataResponse) SetBody(v *UpdateQualityCheckDataResponseBody) *UpdateQualityCheckDataResponse {
	s.Body = v
	return s
}

func (s *UpdateQualityCheckDataResponse) Validate() error {
	return dara.Validate(s)
}
