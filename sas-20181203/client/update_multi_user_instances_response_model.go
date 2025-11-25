// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiUserInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultiUserInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultiUserInstancesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultiUserInstancesResponseBody) *UpdateMultiUserInstancesResponse
	GetBody() *UpdateMultiUserInstancesResponseBody
}

type UpdateMultiUserInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultiUserInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultiUserInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiUserInstancesResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultiUserInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultiUserInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultiUserInstancesResponse) GetBody() *UpdateMultiUserInstancesResponseBody {
	return s.Body
}

func (s *UpdateMultiUserInstancesResponse) SetHeaders(v map[string]*string) *UpdateMultiUserInstancesResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultiUserInstancesResponse) SetStatusCode(v int32) *UpdateMultiUserInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultiUserInstancesResponse) SetBody(v *UpdateMultiUserInstancesResponseBody) *UpdateMultiUserInstancesResponse {
	s.Body = v
	return s
}

func (s *UpdateMultiUserInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
