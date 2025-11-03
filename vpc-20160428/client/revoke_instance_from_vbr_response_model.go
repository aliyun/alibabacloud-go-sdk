// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromVbrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeInstanceFromVbrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeInstanceFromVbrResponse
	GetStatusCode() *int32
	SetBody(v *RevokeInstanceFromVbrResponseBody) *RevokeInstanceFromVbrResponse
	GetBody() *RevokeInstanceFromVbrResponseBody
}

type RevokeInstanceFromVbrResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromVbrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromVbrResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromVbrResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromVbrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeInstanceFromVbrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeInstanceFromVbrResponse) GetBody() *RevokeInstanceFromVbrResponseBody {
	return s.Body
}

func (s *RevokeInstanceFromVbrResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromVbrResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromVbrResponse) SetStatusCode(v int32) *RevokeInstanceFromVbrResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromVbrResponse) SetBody(v *RevokeInstanceFromVbrResponseBody) *RevokeInstanceFromVbrResponse {
	s.Body = v
	return s
}

func (s *RevokeInstanceFromVbrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
