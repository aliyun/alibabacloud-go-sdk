// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOpsOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeOpsOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeOpsOwnerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeOpsOwnerResponseBody) *UpdateNodeOpsOwnerResponse
	GetBody() *UpdateNodeOpsOwnerResponseBody
}

type UpdateNodeOpsOwnerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeOpsOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeOpsOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeOpsOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeOpsOwnerResponse) GetBody() *UpdateNodeOpsOwnerResponseBody {
	return s.Body
}

func (s *UpdateNodeOpsOwnerResponse) SetHeaders(v map[string]*string) *UpdateNodeOpsOwnerResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeOpsOwnerResponse) SetStatusCode(v int32) *UpdateNodeOpsOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeOpsOwnerResponse) SetBody(v *UpdateNodeOpsOwnerResponseBody) *UpdateNodeOpsOwnerResponse {
	s.Body = v
	return s
}

func (s *UpdateNodeOpsOwnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
