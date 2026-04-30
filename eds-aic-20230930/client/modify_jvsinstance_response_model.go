// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJVSInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyJVSInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyJVSInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyJVSInstanceResponseBody) *ModifyJVSInstanceResponse
	GetBody() *ModifyJVSInstanceResponseBody
}

type ModifyJVSInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyJVSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyJVSInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyJVSInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyJVSInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyJVSInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyJVSInstanceResponse) GetBody() *ModifyJVSInstanceResponseBody {
	return s.Body
}

func (s *ModifyJVSInstanceResponse) SetHeaders(v map[string]*string) *ModifyJVSInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyJVSInstanceResponse) SetStatusCode(v int32) *ModifyJVSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyJVSInstanceResponse) SetBody(v *ModifyJVSInstanceResponseBody) *ModifyJVSInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyJVSInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
