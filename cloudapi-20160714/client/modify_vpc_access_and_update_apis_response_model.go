// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAccessAndUpdateApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcAccessAndUpdateApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcAccessAndUpdateApisResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcAccessAndUpdateApisResponseBody) *ModifyVpcAccessAndUpdateApisResponse
	GetBody() *ModifyVpcAccessAndUpdateApisResponseBody
}

type ModifyVpcAccessAndUpdateApisResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcAccessAndUpdateApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcAccessAndUpdateApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAccessAndUpdateApisResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcAccessAndUpdateApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcAccessAndUpdateApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcAccessAndUpdateApisResponse) GetBody() *ModifyVpcAccessAndUpdateApisResponseBody {
	return s.Body
}

func (s *ModifyVpcAccessAndUpdateApisResponse) SetHeaders(v map[string]*string) *ModifyVpcAccessAndUpdateApisResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisResponse) SetStatusCode(v int32) *ModifyVpcAccessAndUpdateApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisResponse) SetBody(v *ModifyVpcAccessAndUpdateApisResponseBody) *ModifyVpcAccessAndUpdateApisResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
