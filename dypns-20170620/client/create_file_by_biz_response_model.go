// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileByBizResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileByBizResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileByBizResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileByBizResponseBody) *CreateFileByBizResponse
	GetBody() *CreateFileByBizResponseBody
}

type CreateFileByBizResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileByBizResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileByBizResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileByBizResponse) GoString() string {
	return s.String()
}

func (s *CreateFileByBizResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileByBizResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileByBizResponse) GetBody() *CreateFileByBizResponseBody {
	return s.Body
}

func (s *CreateFileByBizResponse) SetHeaders(v map[string]*string) *CreateFileByBizResponse {
	s.Headers = v
	return s
}

func (s *CreateFileByBizResponse) SetStatusCode(v int32) *CreateFileByBizResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileByBizResponse) SetBody(v *CreateFileByBizResponseBody) *CreateFileByBizResponse {
	s.Body = v
	return s
}

func (s *CreateFileByBizResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
