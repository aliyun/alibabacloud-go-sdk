// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAppIcpRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmsAppIcpRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmsAppIcpRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmsAppIcpRecordResponseBody) *CreateSmsAppIcpRecordResponse
	GetBody() *CreateSmsAppIcpRecordResponseBody
}

type CreateSmsAppIcpRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsAppIcpRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsAppIcpRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAppIcpRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsAppIcpRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmsAppIcpRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmsAppIcpRecordResponse) GetBody() *CreateSmsAppIcpRecordResponseBody {
	return s.Body
}

func (s *CreateSmsAppIcpRecordResponse) SetHeaders(v map[string]*string) *CreateSmsAppIcpRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsAppIcpRecordResponse) SetStatusCode(v int32) *CreateSmsAppIcpRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponse) SetBody(v *CreateSmsAppIcpRecordResponseBody) *CreateSmsAppIcpRecordResponse {
	s.Body = v
	return s
}

func (s *CreateSmsAppIcpRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
