// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySqlLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySqlLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySqlLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifySqlLogConfigResponseBody) *ModifySqlLogConfigResponse
	GetBody() *ModifySqlLogConfigResponseBody
}

type ModifySqlLogConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySqlLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySqlLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySqlLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySqlLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySqlLogConfigResponse) GetBody() *ModifySqlLogConfigResponseBody {
	return s.Body
}

func (s *ModifySqlLogConfigResponse) SetHeaders(v map[string]*string) *ModifySqlLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySqlLogConfigResponse) SetStatusCode(v int32) *ModifySqlLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySqlLogConfigResponse) SetBody(v *ModifySqlLogConfigResponseBody) *ModifySqlLogConfigResponse {
	s.Body = v
	return s
}

func (s *ModifySqlLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
