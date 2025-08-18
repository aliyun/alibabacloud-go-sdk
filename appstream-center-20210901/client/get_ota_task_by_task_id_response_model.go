// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOtaTaskByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOtaTaskByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOtaTaskByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetOtaTaskByTaskIdResponseBody) *GetOtaTaskByTaskIdResponse
	GetBody() *GetOtaTaskByTaskIdResponseBody
}

type GetOtaTaskByTaskIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOtaTaskByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOtaTaskByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOtaTaskByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOtaTaskByTaskIdResponse) GetBody() *GetOtaTaskByTaskIdResponseBody {
	return s.Body
}

func (s *GetOtaTaskByTaskIdResponse) SetHeaders(v map[string]*string) *GetOtaTaskByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetStatusCode(v int32) *GetOtaTaskByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetBody(v *GetOtaTaskByTaskIdResponseBody) *GetOtaTaskByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) Validate() error {
	return dara.Validate(s)
}
