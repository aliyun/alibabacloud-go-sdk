// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindVbrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindVbrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindVbrResponse
	GetStatusCode() *int32
	SetBody(v *UnbindVbrResponseBody) *UnbindVbrResponse
	GetBody() *UnbindVbrResponseBody
}

type UnbindVbrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindVbrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindVbrResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindVbrResponse) GoString() string {
	return s.String()
}

func (s *UnbindVbrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindVbrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindVbrResponse) GetBody() *UnbindVbrResponseBody {
	return s.Body
}

func (s *UnbindVbrResponse) SetHeaders(v map[string]*string) *UnbindVbrResponse {
	s.Headers = v
	return s
}

func (s *UnbindVbrResponse) SetStatusCode(v int32) *UnbindVbrResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindVbrResponse) SetBody(v *UnbindVbrResponseBody) *UnbindVbrResponse {
	s.Body = v
	return s
}

func (s *UnbindVbrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
