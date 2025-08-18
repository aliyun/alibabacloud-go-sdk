// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineUserInfoResponseBody) *GetRoutineUserInfoResponse
	GetBody() *GetRoutineUserInfoResponseBody
}

type GetRoutineUserInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineUserInfoResponse) GetBody() *GetRoutineUserInfoResponseBody {
	return s.Body
}

func (s *GetRoutineUserInfoResponse) SetHeaders(v map[string]*string) *GetRoutineUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineUserInfoResponse) SetStatusCode(v int32) *GetRoutineUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineUserInfoResponse) SetBody(v *GetRoutineUserInfoResponseBody) *GetRoutineUserInfoResponse {
	s.Body = v
	return s
}

func (s *GetRoutineUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
