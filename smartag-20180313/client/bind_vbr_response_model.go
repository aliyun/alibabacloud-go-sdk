// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVbrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindVbrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindVbrResponse
	GetStatusCode() *int32
	SetBody(v *BindVbrResponseBody) *BindVbrResponse
	GetBody() *BindVbrResponseBody
}

type BindVbrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindVbrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindVbrResponse) String() string {
	return dara.Prettify(s)
}

func (s BindVbrResponse) GoString() string {
	return s.String()
}

func (s *BindVbrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindVbrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindVbrResponse) GetBody() *BindVbrResponseBody {
	return s.Body
}

func (s *BindVbrResponse) SetHeaders(v map[string]*string) *BindVbrResponse {
	s.Headers = v
	return s
}

func (s *BindVbrResponse) SetStatusCode(v int32) *BindVbrResponse {
	s.StatusCode = &v
	return s
}

func (s *BindVbrResponse) SetBody(v *BindVbrResponseBody) *BindVbrResponse {
	s.Body = v
	return s
}

func (s *BindVbrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
