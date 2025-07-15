// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToVbrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantInstanceToVbrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantInstanceToVbrResponse
	GetStatusCode() *int32
	SetBody(v *GrantInstanceToVbrResponseBody) *GrantInstanceToVbrResponse
	GetBody() *GrantInstanceToVbrResponseBody
}

type GrantInstanceToVbrResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToVbrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToVbrResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToVbrResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToVbrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantInstanceToVbrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantInstanceToVbrResponse) GetBody() *GrantInstanceToVbrResponseBody {
	return s.Body
}

func (s *GrantInstanceToVbrResponse) SetHeaders(v map[string]*string) *GrantInstanceToVbrResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToVbrResponse) SetStatusCode(v int32) *GrantInstanceToVbrResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToVbrResponse) SetBody(v *GrantInstanceToVbrResponseBody) *GrantInstanceToVbrResponse {
	s.Body = v
	return s
}

func (s *GrantInstanceToVbrResponse) Validate() error {
	return dara.Validate(s)
}
