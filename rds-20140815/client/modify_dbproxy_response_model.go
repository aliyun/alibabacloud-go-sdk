// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBProxyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBProxyResponseBody) *ModifyDBProxyResponse
	GetBody() *ModifyDBProxyResponseBody
}

type ModifyDBProxyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBProxyResponse) GetBody() *ModifyDBProxyResponseBody {
	return s.Body
}

func (s *ModifyDBProxyResponse) SetHeaders(v map[string]*string) *ModifyDBProxyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBProxyResponse) SetStatusCode(v int32) *ModifyDBProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBProxyResponse) SetBody(v *ModifyDBProxyResponseBody) *ModifyDBProxyResponse {
	s.Body = v
	return s
}

func (s *ModifyDBProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
