// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookInputOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPlaybookInputOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPlaybookInputOutputResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPlaybookInputOutputResponseBody) *ModifyPlaybookInputOutputResponse
	GetBody() *ModifyPlaybookInputOutputResponseBody
}

type ModifyPlaybookInputOutputResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPlaybookInputOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPlaybookInputOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookInputOutputResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPlaybookInputOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPlaybookInputOutputResponse) GetBody() *ModifyPlaybookInputOutputResponseBody {
	return s.Body
}

func (s *ModifyPlaybookInputOutputResponse) SetHeaders(v map[string]*string) *ModifyPlaybookInputOutputResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) SetStatusCode(v int32) *ModifyPlaybookInputOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) SetBody(v *ModifyPlaybookInputOutputResponseBody) *ModifyPlaybookInputOutputResponse {
	s.Body = v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
