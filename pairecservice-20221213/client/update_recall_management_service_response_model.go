// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecallManagementServiceResponseBody) *UpdateRecallManagementServiceResponse
	GetBody() *UpdateRecallManagementServiceResponseBody
}

type UpdateRecallManagementServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecallManagementServiceResponse) GetBody() *UpdateRecallManagementServiceResponseBody {
	return s.Body
}

func (s *UpdateRecallManagementServiceResponse) SetHeaders(v map[string]*string) *UpdateRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecallManagementServiceResponse) SetStatusCode(v int32) *UpdateRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecallManagementServiceResponse) SetBody(v *UpdateRecallManagementServiceResponseBody) *UpdateRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
