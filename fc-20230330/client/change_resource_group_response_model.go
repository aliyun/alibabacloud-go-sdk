// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeResourceGroupOutput) *ChangeResourceGroupResponse
	GetBody() *ChangeResourceGroupOutput
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string         `json:"headers" xml:"headers"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeResourceGroupResponse) GetBody() *ChangeResourceGroupOutput {
	return s.Body
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupOutput) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
