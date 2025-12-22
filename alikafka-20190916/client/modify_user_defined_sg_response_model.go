// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserDefinedSgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserDefinedSgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserDefinedSgResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserDefinedSgResponseBody) *ModifyUserDefinedSgResponse
	GetBody() *ModifyUserDefinedSgResponseBody
}

type ModifyUserDefinedSgResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserDefinedSgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserDefinedSgResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserDefinedSgResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserDefinedSgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserDefinedSgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserDefinedSgResponse) GetBody() *ModifyUserDefinedSgResponseBody {
	return s.Body
}

func (s *ModifyUserDefinedSgResponse) SetHeaders(v map[string]*string) *ModifyUserDefinedSgResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserDefinedSgResponse) SetStatusCode(v int32) *ModifyUserDefinedSgResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserDefinedSgResponse) SetBody(v *ModifyUserDefinedSgResponseBody) *ModifyUserDefinedSgResponse {
	s.Body = v
	return s
}

func (s *ModifyUserDefinedSgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
