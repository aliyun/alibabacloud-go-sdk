// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindDBResourceGroupWithUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindDBResourceGroupWithUserResponse
	GetStatusCode() *int32
	SetBody(v *BindDBResourceGroupWithUserResponseBody) *BindDBResourceGroupWithUserResponse
	GetBody() *BindDBResourceGroupWithUserResponseBody
}

type BindDBResourceGroupWithUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindDBResourceGroupWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindDBResourceGroupWithUserResponse) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithUserResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindDBResourceGroupWithUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindDBResourceGroupWithUserResponse) GetBody() *BindDBResourceGroupWithUserResponseBody {
	return s.Body
}

func (s *BindDBResourceGroupWithUserResponse) SetHeaders(v map[string]*string) *BindDBResourceGroupWithUserResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) SetStatusCode(v int32) *BindDBResourceGroupWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) SetBody(v *BindDBResourceGroupWithUserResponseBody) *BindDBResourceGroupWithUserResponse {
	s.Body = v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
