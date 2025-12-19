// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationSpecResponseBody) *ModifyApplicationSpecResponse
	GetBody() *ModifyApplicationSpecResponseBody
}

type ModifyApplicationSpecResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationSpecResponse) GetBody() *ModifyApplicationSpecResponseBody {
	return s.Body
}

func (s *ModifyApplicationSpecResponse) SetHeaders(v map[string]*string) *ModifyApplicationSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationSpecResponse) SetStatusCode(v int32) *ModifyApplicationSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationSpecResponse) SetBody(v *ModifyApplicationSpecResponseBody) *ModifyApplicationSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
