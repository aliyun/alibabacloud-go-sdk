// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecallManagementConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecallManagementConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecallManagementConfigResponseBody) *UpdateRecallManagementConfigResponse
	GetBody() *UpdateRecallManagementConfigResponseBody
}

type UpdateRecallManagementConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecallManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecallManagementConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecallManagementConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecallManagementConfigResponse) GetBody() *UpdateRecallManagementConfigResponseBody {
	return s.Body
}

func (s *UpdateRecallManagementConfigResponse) SetHeaders(v map[string]*string) *UpdateRecallManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecallManagementConfigResponse) SetStatusCode(v int32) *UpdateRecallManagementConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecallManagementConfigResponse) SetBody(v *UpdateRecallManagementConfigResponseBody) *UpdateRecallManagementConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateRecallManagementConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
