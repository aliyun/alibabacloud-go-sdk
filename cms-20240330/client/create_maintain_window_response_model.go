// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaintainWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMaintainWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMaintainWindowResponse
	GetStatusCode() *int32
	SetBody(v *CreateMaintainWindowResponseBody) *CreateMaintainWindowResponse
	GetBody() *CreateMaintainWindowResponseBody
}

type CreateMaintainWindowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMaintainWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMaintainWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMaintainWindowResponse) GoString() string {
	return s.String()
}

func (s *CreateMaintainWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMaintainWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMaintainWindowResponse) GetBody() *CreateMaintainWindowResponseBody {
	return s.Body
}

func (s *CreateMaintainWindowResponse) SetHeaders(v map[string]*string) *CreateMaintainWindowResponse {
	s.Headers = v
	return s
}

func (s *CreateMaintainWindowResponse) SetStatusCode(v int32) *CreateMaintainWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMaintainWindowResponse) SetBody(v *CreateMaintainWindowResponseBody) *CreateMaintainWindowResponse {
	s.Body = v
	return s
}

func (s *CreateMaintainWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
