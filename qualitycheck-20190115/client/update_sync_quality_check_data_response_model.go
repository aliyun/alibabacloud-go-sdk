// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSyncQualityCheckDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSyncQualityCheckDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSyncQualityCheckDataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSyncQualityCheckDataResponseBody) *UpdateSyncQualityCheckDataResponse
	GetBody() *UpdateSyncQualityCheckDataResponseBody
}

type UpdateSyncQualityCheckDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSyncQualityCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSyncQualityCheckDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSyncQualityCheckDataResponse) GetBody() *UpdateSyncQualityCheckDataResponseBody {
	return s.Body
}

func (s *UpdateSyncQualityCheckDataResponse) SetHeaders(v map[string]*string) *UpdateSyncQualityCheckDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) SetStatusCode(v int32) *UpdateSyncQualityCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) SetBody(v *UpdateSyncQualityCheckDataResponseBody) *UpdateSyncQualityCheckDataResponse {
	s.Body = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) Validate() error {
	return dara.Validate(s)
}
