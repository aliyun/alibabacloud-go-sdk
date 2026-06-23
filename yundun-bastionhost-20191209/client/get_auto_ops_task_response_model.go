// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoOpsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoOpsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoOpsTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoOpsTaskResponseBody) *GetAutoOpsTaskResponse
	GetBody() *GetAutoOpsTaskResponseBody
}

type GetAutoOpsTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoOpsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoOpsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoOpsTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAutoOpsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoOpsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoOpsTaskResponse) GetBody() *GetAutoOpsTaskResponseBody {
	return s.Body
}

func (s *GetAutoOpsTaskResponse) SetHeaders(v map[string]*string) *GetAutoOpsTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAutoOpsTaskResponse) SetStatusCode(v int32) *GetAutoOpsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoOpsTaskResponse) SetBody(v *GetAutoOpsTaskResponseBody) *GetAutoOpsTaskResponse {
	s.Body = v
	return s
}

func (s *GetAutoOpsTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
