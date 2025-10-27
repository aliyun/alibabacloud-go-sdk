// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateVulsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateVulsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateVulsResponse
	GetStatusCode() *int32
	SetBody(v *OperateVulsResponseBody) *OperateVulsResponse
	GetBody() *OperateVulsResponseBody
}

type OperateVulsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateVulsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateVulsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateVulsResponse) GoString() string {
	return s.String()
}

func (s *OperateVulsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateVulsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateVulsResponse) GetBody() *OperateVulsResponseBody {
	return s.Body
}

func (s *OperateVulsResponse) SetHeaders(v map[string]*string) *OperateVulsResponse {
	s.Headers = v
	return s
}

func (s *OperateVulsResponse) SetStatusCode(v int32) *OperateVulsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateVulsResponse) SetBody(v *OperateVulsResponseBody) *OperateVulsResponse {
	s.Body = v
	return s
}

func (s *OperateVulsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
