// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceActiveOpsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceActiveOpsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceActiveOpsTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceActiveOpsTaskResponseBody) *CreateInstanceActiveOpsTaskResponse
	GetBody() *CreateInstanceActiveOpsTaskResponseBody
}

type CreateInstanceActiveOpsTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceActiveOpsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceActiveOpsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceActiveOpsTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceActiveOpsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceActiveOpsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceActiveOpsTaskResponse) GetBody() *CreateInstanceActiveOpsTaskResponseBody {
	return s.Body
}

func (s *CreateInstanceActiveOpsTaskResponse) SetHeaders(v map[string]*string) *CreateInstanceActiveOpsTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceActiveOpsTaskResponse) SetStatusCode(v int32) *CreateInstanceActiveOpsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceActiveOpsTaskResponse) SetBody(v *CreateInstanceActiveOpsTaskResponseBody) *CreateInstanceActiveOpsTaskResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceActiveOpsTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
