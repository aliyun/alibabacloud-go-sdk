// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtcAsrTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRtcAsrTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRtcAsrTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRtcAsrTaskResponseBody) *CreateRtcAsrTaskResponse
	GetBody() *CreateRtcAsrTaskResponseBody
}

type CreateRtcAsrTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRtcAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRtcAsrTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRtcAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRtcAsrTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRtcAsrTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRtcAsrTaskResponse) GetBody() *CreateRtcAsrTaskResponseBody {
	return s.Body
}

func (s *CreateRtcAsrTaskResponse) SetHeaders(v map[string]*string) *CreateRtcAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRtcAsrTaskResponse) SetStatusCode(v int32) *CreateRtcAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRtcAsrTaskResponse) SetBody(v *CreateRtcAsrTaskResponseBody) *CreateRtcAsrTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRtcAsrTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
