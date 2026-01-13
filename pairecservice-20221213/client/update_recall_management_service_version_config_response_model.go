// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementServiceVersionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecallManagementServiceVersionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecallManagementServiceVersionConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecallManagementServiceVersionConfigResponseBody) *UpdateRecallManagementServiceVersionConfigResponse
	GetBody() *UpdateRecallManagementServiceVersionConfigResponseBody
}

type UpdateRecallManagementServiceVersionConfigResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecallManagementServiceVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) GetBody() *UpdateRecallManagementServiceVersionConfigResponseBody {
	return s.Body
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) SetHeaders(v map[string]*string) *UpdateRecallManagementServiceVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) SetStatusCode(v int32) *UpdateRecallManagementServiceVersionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) SetBody(v *UpdateRecallManagementServiceVersionConfigResponseBody) *UpdateRecallManagementServiceVersionConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
