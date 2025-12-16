// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateABTestGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateABTestGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateABTestGroupResponseBody) *CreateABTestGroupResponse
	GetBody() *CreateABTestGroupResponseBody
}

type CreateABTestGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABTestGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateABTestGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateABTestGroupResponse) GetBody() *CreateABTestGroupResponseBody {
	return s.Body
}

func (s *CreateABTestGroupResponse) SetHeaders(v map[string]*string) *CreateABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestGroupResponse) SetStatusCode(v int32) *CreateABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABTestGroupResponse) SetBody(v *CreateABTestGroupResponseBody) *CreateABTestGroupResponse {
	s.Body = v
	return s
}

func (s *CreateABTestGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
