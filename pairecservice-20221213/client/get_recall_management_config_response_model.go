// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecallManagementConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecallManagementConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetRecallManagementConfigResponseBody) *GetRecallManagementConfigResponse
	GetBody() *GetRecallManagementConfigResponseBody
}

type GetRecallManagementConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecallManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecallManagementConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *GetRecallManagementConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecallManagementConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecallManagementConfigResponse) GetBody() *GetRecallManagementConfigResponseBody {
	return s.Body
}

func (s *GetRecallManagementConfigResponse) SetHeaders(v map[string]*string) *GetRecallManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *GetRecallManagementConfigResponse) SetStatusCode(v int32) *GetRecallManagementConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecallManagementConfigResponse) SetBody(v *GetRecallManagementConfigResponseBody) *GetRecallManagementConfigResponse {
	s.Body = v
	return s
}

func (s *GetRecallManagementConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
