// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmonMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmonMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *GetEmonMonitorDataResponseBody) *GetEmonMonitorDataResponse
	GetBody() *GetEmonMonitorDataResponseBody
}

type GetEmonMonitorDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmonMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmonMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmonMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmonMonitorDataResponse) GetBody() *GetEmonMonitorDataResponseBody {
	return s.Body
}

func (s *GetEmonMonitorDataResponse) SetHeaders(v map[string]*string) *GetEmonMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *GetEmonMonitorDataResponse) SetStatusCode(v int32) *GetEmonMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmonMonitorDataResponse) SetBody(v *GetEmonMonitorDataResponseBody) *GetEmonMonitorDataResponse {
	s.Body = v
	return s
}

func (s *GetEmonMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
