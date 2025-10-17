// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGadInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGadInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGadInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGadInstanceNameResponseBody) *ModifyGadInstanceNameResponse
	GetBody() *ModifyGadInstanceNameResponseBody
}

type ModifyGadInstanceNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGadInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGadInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGadInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyGadInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGadInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGadInstanceNameResponse) GetBody() *ModifyGadInstanceNameResponseBody {
	return s.Body
}

func (s *ModifyGadInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyGadInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyGadInstanceNameResponse) SetStatusCode(v int32) *ModifyGadInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGadInstanceNameResponse) SetBody(v *ModifyGadInstanceNameResponseBody) *ModifyGadInstanceNameResponse {
	s.Body = v
	return s
}

func (s *ModifyGadInstanceNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
