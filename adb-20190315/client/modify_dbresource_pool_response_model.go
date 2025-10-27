// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourcePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBResourcePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBResourcePoolResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBResourcePoolResponseBody) *ModifyDBResourcePoolResponse
	GetBody() *ModifyDBResourcePoolResponseBody
}

type ModifyDBResourcePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBResourcePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBResourcePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBResourcePoolResponse) GetBody() *ModifyDBResourcePoolResponseBody {
	return s.Body
}

func (s *ModifyDBResourcePoolResponse) SetHeaders(v map[string]*string) *ModifyDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBResourcePoolResponse) SetStatusCode(v int32) *ModifyDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBResourcePoolResponse) SetBody(v *ModifyDBResourcePoolResponseBody) *ModifyDBResourcePoolResponse {
	s.Body = v
	return s
}

func (s *ModifyDBResourcePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
