// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagWithUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTagWithUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTagWithUuidResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTagWithUuidResponseBody) *ModifyTagWithUuidResponse
	GetBody() *ModifyTagWithUuidResponseBody
}

type ModifyTagWithUuidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTagWithUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTagWithUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagWithUuidResponse) GoString() string {
	return s.String()
}

func (s *ModifyTagWithUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTagWithUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTagWithUuidResponse) GetBody() *ModifyTagWithUuidResponseBody {
	return s.Body
}

func (s *ModifyTagWithUuidResponse) SetHeaders(v map[string]*string) *ModifyTagWithUuidResponse {
	s.Headers = v
	return s
}

func (s *ModifyTagWithUuidResponse) SetStatusCode(v int32) *ModifyTagWithUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTagWithUuidResponse) SetBody(v *ModifyTagWithUuidResponseBody) *ModifyTagWithUuidResponse {
	s.Body = v
	return s
}

func (s *ModifyTagWithUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
