// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCasterComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCasterComponentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCasterComponentResponseBody) *ModifyCasterComponentResponse
	GetBody() *ModifyCasterComponentResponseBody
}

type ModifyCasterComponentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCasterComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCasterComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterComponentResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCasterComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCasterComponentResponse) GetBody() *ModifyCasterComponentResponseBody {
	return s.Body
}

func (s *ModifyCasterComponentResponse) SetHeaders(v map[string]*string) *ModifyCasterComponentResponse {
	s.Headers = v
	return s
}

func (s *ModifyCasterComponentResponse) SetStatusCode(v int32) *ModifyCasterComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCasterComponentResponse) SetBody(v *ModifyCasterComponentResponseBody) *ModifyCasterComponentResponse {
	s.Body = v
	return s
}

func (s *ModifyCasterComponentResponse) Validate() error {
	return dara.Validate(s)
}
