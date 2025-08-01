// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineEvalTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOnlineEvalTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOnlineEvalTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetOnlineEvalTaskResponseBody) *GetOnlineEvalTaskResponse
	GetBody() *GetOnlineEvalTaskResponseBody
}

type GetOnlineEvalTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnlineEvalTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineEvalTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOnlineEvalTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOnlineEvalTaskResponse) GetBody() *GetOnlineEvalTaskResponseBody {
	return s.Body
}

func (s *GetOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *GetOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineEvalTaskResponse) SetStatusCode(v int32) *GetOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnlineEvalTaskResponse) SetBody(v *GetOnlineEvalTaskResponseBody) *GetOnlineEvalTaskResponse {
	s.Body = v
	return s
}

func (s *GetOnlineEvalTaskResponse) Validate() error {
	return dara.Validate(s)
}
