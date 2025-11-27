// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCNodePoolResponseBody) *CreateRCNodePoolResponse
	GetBody() *CreateRCNodePoolResponseBody
}

type CreateRCNodePoolResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCNodePoolResponse) GetBody() *CreateRCNodePoolResponseBody {
	return s.Body
}

func (s *CreateRCNodePoolResponse) SetHeaders(v map[string]*string) *CreateRCNodePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateRCNodePoolResponse) SetStatusCode(v int32) *CreateRCNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCNodePoolResponse) SetBody(v *CreateRCNodePoolResponseBody) *CreateRCNodePoolResponse {
	s.Body = v
	return s
}

func (s *CreateRCNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
