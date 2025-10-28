// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationBaseInfoResponseBody) *UpdateApplicationBaseInfoResponse
	GetBody() *UpdateApplicationBaseInfoResponseBody
}

type UpdateApplicationBaseInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationBaseInfoResponse) GetBody() *UpdateApplicationBaseInfoResponseBody {
	return s.Body
}

func (s *UpdateApplicationBaseInfoResponse) SetHeaders(v map[string]*string) *UpdateApplicationBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationBaseInfoResponse) SetStatusCode(v int32) *UpdateApplicationBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponse) SetBody(v *UpdateApplicationBaseInfoResponseBody) *UpdateApplicationBaseInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
