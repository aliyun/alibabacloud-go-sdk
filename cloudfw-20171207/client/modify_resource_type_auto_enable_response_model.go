// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceTypeAutoEnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceTypeAutoEnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceTypeAutoEnableResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceTypeAutoEnableResponseBody) *ModifyResourceTypeAutoEnableResponse
	GetBody() *ModifyResourceTypeAutoEnableResponseBody
}

type ModifyResourceTypeAutoEnableResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceTypeAutoEnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceTypeAutoEnableResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceTypeAutoEnableResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceTypeAutoEnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceTypeAutoEnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceTypeAutoEnableResponse) GetBody() *ModifyResourceTypeAutoEnableResponseBody {
	return s.Body
}

func (s *ModifyResourceTypeAutoEnableResponse) SetHeaders(v map[string]*string) *ModifyResourceTypeAutoEnableResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceTypeAutoEnableResponse) SetStatusCode(v int32) *ModifyResourceTypeAutoEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceTypeAutoEnableResponse) SetBody(v *ModifyResourceTypeAutoEnableResponseBody) *ModifyResourceTypeAutoEnableResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceTypeAutoEnableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
