// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumOcuStatisticDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumOcuStatisticDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumOcuStatisticDataResponse
	GetStatusCode() *int32
	SetBody(v *GetRumOcuStatisticDataResponseBody) *GetRumOcuStatisticDataResponse
	GetBody() *GetRumOcuStatisticDataResponseBody
}

type GetRumOcuStatisticDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumOcuStatisticDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumOcuStatisticDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumOcuStatisticDataResponse) GoString() string {
	return s.String()
}

func (s *GetRumOcuStatisticDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumOcuStatisticDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumOcuStatisticDataResponse) GetBody() *GetRumOcuStatisticDataResponseBody {
	return s.Body
}

func (s *GetRumOcuStatisticDataResponse) SetHeaders(v map[string]*string) *GetRumOcuStatisticDataResponse {
	s.Headers = v
	return s
}

func (s *GetRumOcuStatisticDataResponse) SetStatusCode(v int32) *GetRumOcuStatisticDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumOcuStatisticDataResponse) SetBody(v *GetRumOcuStatisticDataResponseBody) *GetRumOcuStatisticDataResponse {
	s.Body = v
	return s
}

func (s *GetRumOcuStatisticDataResponse) Validate() error {
	return dara.Validate(s)
}
