// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSagInstanceToVbrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantSagInstanceToVbrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantSagInstanceToVbrResponse
	GetStatusCode() *int32
	SetBody(v *GrantSagInstanceToVbrResponseBody) *GrantSagInstanceToVbrResponse
	GetBody() *GrantSagInstanceToVbrResponseBody
}

type GrantSagInstanceToVbrResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantSagInstanceToVbrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantSagInstanceToVbrResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantSagInstanceToVbrResponse) GoString() string {
	return s.String()
}

func (s *GrantSagInstanceToVbrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantSagInstanceToVbrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantSagInstanceToVbrResponse) GetBody() *GrantSagInstanceToVbrResponseBody {
	return s.Body
}

func (s *GrantSagInstanceToVbrResponse) SetHeaders(v map[string]*string) *GrantSagInstanceToVbrResponse {
	s.Headers = v
	return s
}

func (s *GrantSagInstanceToVbrResponse) SetStatusCode(v int32) *GrantSagInstanceToVbrResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantSagInstanceToVbrResponse) SetBody(v *GrantSagInstanceToVbrResponseBody) *GrantSagInstanceToVbrResponse {
	s.Body = v
	return s
}

func (s *GrantSagInstanceToVbrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
