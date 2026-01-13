// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetRecallManagementServiceResponseBody) *GetRecallManagementServiceResponse
	GetBody() *GetRecallManagementServiceResponseBody
}

type GetRecallManagementServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecallManagementServiceResponse) GetBody() *GetRecallManagementServiceResponseBody {
	return s.Body
}

func (s *GetRecallManagementServiceResponse) SetHeaders(v map[string]*string) *GetRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *GetRecallManagementServiceResponse) SetStatusCode(v int32) *GetRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecallManagementServiceResponse) SetBody(v *GetRecallManagementServiceResponseBody) *GetRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *GetRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
