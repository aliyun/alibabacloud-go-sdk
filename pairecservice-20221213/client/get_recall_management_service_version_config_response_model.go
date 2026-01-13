// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceVersionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecallManagementServiceVersionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecallManagementServiceVersionConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetRecallManagementServiceVersionConfigResponseBody) *GetRecallManagementServiceVersionConfigResponse
	GetBody() *GetRecallManagementServiceVersionConfigResponseBody
}

type GetRecallManagementServiceVersionConfigResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecallManagementServiceVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecallManagementServiceVersionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecallManagementServiceVersionConfigResponse) GetBody() *GetRecallManagementServiceVersionConfigResponseBody {
	return s.Body
}

func (s *GetRecallManagementServiceVersionConfigResponse) SetHeaders(v map[string]*string) *GetRecallManagementServiceVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponse) SetStatusCode(v int32) *GetRecallManagementServiceVersionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponse) SetBody(v *GetRecallManagementServiceVersionConfigResponseBody) *GetRecallManagementServiceVersionConfigResponse {
	s.Body = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
