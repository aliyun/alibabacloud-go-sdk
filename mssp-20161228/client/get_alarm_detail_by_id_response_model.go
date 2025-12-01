// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmDetailByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlarmDetailByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlarmDetailByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetAlarmDetailByIdResponseBody) *GetAlarmDetailByIdResponse
	GetBody() *GetAlarmDetailByIdResponseBody
}

type GetAlarmDetailByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlarmDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlarmDetailByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlarmDetailByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlarmDetailByIdResponse) GetBody() *GetAlarmDetailByIdResponseBody {
	return s.Body
}

func (s *GetAlarmDetailByIdResponse) SetHeaders(v map[string]*string) *GetAlarmDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAlarmDetailByIdResponse) SetStatusCode(v int32) *GetAlarmDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlarmDetailByIdResponse) SetBody(v *GetAlarmDetailByIdResponseBody) *GetAlarmDetailByIdResponse {
	s.Body = v
	return s
}

func (s *GetAlarmDetailByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
