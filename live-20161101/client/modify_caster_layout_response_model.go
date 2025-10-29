// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCasterLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCasterLayoutResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCasterLayoutResponseBody) *ModifyCasterLayoutResponse
	GetBody() *ModifyCasterLayoutResponseBody
}

type ModifyCasterLayoutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCasterLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCasterLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterLayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCasterLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCasterLayoutResponse) GetBody() *ModifyCasterLayoutResponseBody {
	return s.Body
}

func (s *ModifyCasterLayoutResponse) SetHeaders(v map[string]*string) *ModifyCasterLayoutResponse {
	s.Headers = v
	return s
}

func (s *ModifyCasterLayoutResponse) SetStatusCode(v int32) *ModifyCasterLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCasterLayoutResponse) SetBody(v *ModifyCasterLayoutResponseBody) *ModifyCasterLayoutResponse {
	s.Body = v
	return s
}

func (s *ModifyCasterLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
