// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *OfflineRecallManagementServiceResponseBody) *OfflineRecallManagementServiceResponse
	GetBody() *OfflineRecallManagementServiceResponseBody
}

type OfflineRecallManagementServiceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *OfflineRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineRecallManagementServiceResponse) GetBody() *OfflineRecallManagementServiceResponseBody {
	return s.Body
}

func (s *OfflineRecallManagementServiceResponse) SetHeaders(v map[string]*string) *OfflineRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *OfflineRecallManagementServiceResponse) SetStatusCode(v int32) *OfflineRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineRecallManagementServiceResponse) SetBody(v *OfflineRecallManagementServiceResponseBody) *OfflineRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *OfflineRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
