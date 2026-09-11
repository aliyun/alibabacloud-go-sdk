// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintainWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMaintainWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMaintainWindowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMaintainWindowResponseBody) *UpdateMaintainWindowResponse
	GetBody() *UpdateMaintainWindowResponseBody
}

type UpdateMaintainWindowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMaintainWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMaintainWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintainWindowResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaintainWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMaintainWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMaintainWindowResponse) GetBody() *UpdateMaintainWindowResponseBody {
	return s.Body
}

func (s *UpdateMaintainWindowResponse) SetHeaders(v map[string]*string) *UpdateMaintainWindowResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaintainWindowResponse) SetStatusCode(v int32) *UpdateMaintainWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMaintainWindowResponse) SetBody(v *UpdateMaintainWindowResponseBody) *UpdateMaintainWindowResponse {
	s.Body = v
	return s
}

func (s *UpdateMaintainWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
