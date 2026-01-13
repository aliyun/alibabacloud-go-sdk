// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecallManagementConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecallManagementConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecallManagementConfigResponseBody) *CreateRecallManagementConfigResponse
	GetBody() *CreateRecallManagementConfigResponseBody
}

type CreateRecallManagementConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecallManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecallManagementConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecallManagementConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecallManagementConfigResponse) GetBody() *CreateRecallManagementConfigResponseBody {
	return s.Body
}

func (s *CreateRecallManagementConfigResponse) SetHeaders(v map[string]*string) *CreateRecallManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateRecallManagementConfigResponse) SetStatusCode(v int32) *CreateRecallManagementConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecallManagementConfigResponse) SetBody(v *CreateRecallManagementConfigResponseBody) *CreateRecallManagementConfigResponse {
	s.Body = v
	return s
}

func (s *CreateRecallManagementConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
