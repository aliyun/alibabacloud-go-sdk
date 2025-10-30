// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFieldResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFieldResponseBody) *ModifyFieldResponse
	GetBody() *ModifyFieldResponseBody
}

type ModifyFieldResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFieldResponse) GoString() string {
	return s.String()
}

func (s *ModifyFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFieldResponse) GetBody() *ModifyFieldResponseBody {
	return s.Body
}

func (s *ModifyFieldResponse) SetHeaders(v map[string]*string) *ModifyFieldResponse {
	s.Headers = v
	return s
}

func (s *ModifyFieldResponse) SetStatusCode(v int32) *ModifyFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFieldResponse) SetBody(v *ModifyFieldResponseBody) *ModifyFieldResponse {
	s.Body = v
	return s
}

func (s *ModifyFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
