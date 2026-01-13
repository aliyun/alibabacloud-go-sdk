// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceVersionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecallManagementServiceVersionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecallManagementServiceVersionConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecallManagementServiceVersionConfigResponseBody) *CreateRecallManagementServiceVersionConfigResponse
	GetBody() *CreateRecallManagementServiceVersionConfigResponseBody
}

type CreateRecallManagementServiceVersionConfigResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecallManagementServiceVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecallManagementServiceVersionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecallManagementServiceVersionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecallManagementServiceVersionConfigResponse) GetBody() *CreateRecallManagementServiceVersionConfigResponseBody {
	return s.Body
}

func (s *CreateRecallManagementServiceVersionConfigResponse) SetHeaders(v map[string]*string) *CreateRecallManagementServiceVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigResponse) SetStatusCode(v int32) *CreateRecallManagementServiceVersionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigResponse) SetBody(v *CreateRecallManagementServiceVersionConfigResponseBody) *CreateRecallManagementServiceVersionConfigResponse {
	s.Body = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
