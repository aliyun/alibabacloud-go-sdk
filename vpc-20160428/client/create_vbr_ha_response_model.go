// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVbrHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVbrHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVbrHaResponse
	GetStatusCode() *int32
	SetBody(v *CreateVbrHaResponseBody) *CreateVbrHaResponse
	GetBody() *CreateVbrHaResponseBody
}

type CreateVbrHaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVbrHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVbrHaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVbrHaResponse) GoString() string {
	return s.String()
}

func (s *CreateVbrHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVbrHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVbrHaResponse) GetBody() *CreateVbrHaResponseBody {
	return s.Body
}

func (s *CreateVbrHaResponse) SetHeaders(v map[string]*string) *CreateVbrHaResponse {
	s.Headers = v
	return s
}

func (s *CreateVbrHaResponse) SetStatusCode(v int32) *CreateVbrHaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVbrHaResponse) SetBody(v *CreateVbrHaResponseBody) *CreateVbrHaResponse {
	s.Body = v
	return s
}

func (s *CreateVbrHaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
