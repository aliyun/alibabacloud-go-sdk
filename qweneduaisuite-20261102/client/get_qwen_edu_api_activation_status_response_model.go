// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQwenEduApiActivationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQwenEduApiActivationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQwenEduApiActivationStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetQwenEduApiActivationStatusResponseBody) *GetQwenEduApiActivationStatusResponse
	GetBody() *GetQwenEduApiActivationStatusResponseBody
}

type GetQwenEduApiActivationStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQwenEduApiActivationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQwenEduApiActivationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQwenEduApiActivationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetQwenEduApiActivationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQwenEduApiActivationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQwenEduApiActivationStatusResponse) GetBody() *GetQwenEduApiActivationStatusResponseBody {
	return s.Body
}

func (s *GetQwenEduApiActivationStatusResponse) SetHeaders(v map[string]*string) *GetQwenEduApiActivationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetQwenEduApiActivationStatusResponse) SetStatusCode(v int32) *GetQwenEduApiActivationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponse) SetBody(v *GetQwenEduApiActivationStatusResponseBody) *GetQwenEduApiActivationStatusResponse {
	s.Body = v
	return s
}

func (s *GetQwenEduApiActivationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
