// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySpecTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySpecTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySpecTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySpecTypeResponseBody) *ModifySpecTypeResponse
	GetBody() *ModifySpecTypeResponseBody
}

type ModifySpecTypeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySpecTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySpecTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifySpecTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySpecTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySpecTypeResponse) GetBody() *ModifySpecTypeResponseBody {
	return s.Body
}

func (s *ModifySpecTypeResponse) SetHeaders(v map[string]*string) *ModifySpecTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifySpecTypeResponse) SetStatusCode(v int32) *ModifySpecTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySpecTypeResponse) SetBody(v *ModifySpecTypeResponseBody) *ModifySpecTypeResponse {
	s.Body = v
	return s
}

func (s *ModifySpecTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
