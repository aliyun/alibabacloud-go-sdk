// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnceTaskResultInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOnceTaskResultInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOnceTaskResultInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetOnceTaskResultInfoResponseBody) *GetOnceTaskResultInfoResponse
	GetBody() *GetOnceTaskResultInfoResponseBody
}

type GetOnceTaskResultInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnceTaskResultInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnceTaskResultInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOnceTaskResultInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOnceTaskResultInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOnceTaskResultInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOnceTaskResultInfoResponse) GetBody() *GetOnceTaskResultInfoResponseBody {
	return s.Body
}

func (s *GetOnceTaskResultInfoResponse) SetHeaders(v map[string]*string) *GetOnceTaskResultInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOnceTaskResultInfoResponse) SetStatusCode(v int32) *GetOnceTaskResultInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnceTaskResultInfoResponse) SetBody(v *GetOnceTaskResultInfoResponseBody) *GetOnceTaskResultInfoResponse {
	s.Body = v
	return s
}

func (s *GetOnceTaskResultInfoResponse) Validate() error {
	return dara.Validate(s)
}
