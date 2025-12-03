// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHasRootPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHasRootPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHasRootPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHasRootPasswordResponseBody) *ModifyHasRootPasswordResponse
	GetBody() *ModifyHasRootPasswordResponseBody
}

type ModifyHasRootPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHasRootPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHasRootPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHasRootPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyHasRootPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHasRootPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHasRootPasswordResponse) GetBody() *ModifyHasRootPasswordResponseBody {
	return s.Body
}

func (s *ModifyHasRootPasswordResponse) SetHeaders(v map[string]*string) *ModifyHasRootPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyHasRootPasswordResponse) SetStatusCode(v int32) *ModifyHasRootPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHasRootPasswordResponse) SetBody(v *ModifyHasRootPasswordResponseBody) *ModifyHasRootPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyHasRootPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
