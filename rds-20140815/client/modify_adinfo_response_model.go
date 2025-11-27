// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyADInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyADInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyADInfoResponseBody) *ModifyADInfoResponse
	GetBody() *ModifyADInfoResponseBody
}

type ModifyADInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyADInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyADInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyADInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyADInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyADInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyADInfoResponse) GetBody() *ModifyADInfoResponseBody {
	return s.Body
}

func (s *ModifyADInfoResponse) SetHeaders(v map[string]*string) *ModifyADInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyADInfoResponse) SetStatusCode(v int32) *ModifyADInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyADInfoResponse) SetBody(v *ModifyADInfoResponseBody) *ModifyADInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyADInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
