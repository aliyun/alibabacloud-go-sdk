// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecallManagementServiceResponseBody) *CreateRecallManagementServiceResponse
	GetBody() *CreateRecallManagementServiceResponseBody
}

type CreateRecallManagementServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecallManagementServiceResponse) GetBody() *CreateRecallManagementServiceResponseBody {
	return s.Body
}

func (s *CreateRecallManagementServiceResponse) SetHeaders(v map[string]*string) *CreateRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateRecallManagementServiceResponse) SetStatusCode(v int32) *CreateRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecallManagementServiceResponse) SetBody(v *CreateRecallManagementServiceResponseBody) *CreateRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *CreateRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
