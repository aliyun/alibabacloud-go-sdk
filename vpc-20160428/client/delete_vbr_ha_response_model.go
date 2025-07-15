// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVbrHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVbrHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVbrHaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVbrHaResponseBody) *DeleteVbrHaResponse
	GetBody() *DeleteVbrHaResponseBody
}

type DeleteVbrHaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVbrHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVbrHaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVbrHaResponse) GoString() string {
	return s.String()
}

func (s *DeleteVbrHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVbrHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVbrHaResponse) GetBody() *DeleteVbrHaResponseBody {
	return s.Body
}

func (s *DeleteVbrHaResponse) SetHeaders(v map[string]*string) *DeleteVbrHaResponse {
	s.Headers = v
	return s
}

func (s *DeleteVbrHaResponse) SetStatusCode(v int32) *DeleteVbrHaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVbrHaResponse) SetBody(v *DeleteVbrHaResponseBody) *DeleteVbrHaResponse {
	s.Body = v
	return s
}

func (s *DeleteVbrHaResponse) Validate() error {
	return dara.Validate(s)
}
