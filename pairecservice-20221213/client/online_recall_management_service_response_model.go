// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *OnlineRecallManagementServiceResponseBody) *OnlineRecallManagementServiceResponse
	GetBody() *OnlineRecallManagementServiceResponseBody
}

type OnlineRecallManagementServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *OnlineRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineRecallManagementServiceResponse) GetBody() *OnlineRecallManagementServiceResponseBody {
	return s.Body
}

func (s *OnlineRecallManagementServiceResponse) SetHeaders(v map[string]*string) *OnlineRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *OnlineRecallManagementServiceResponse) SetStatusCode(v int32) *OnlineRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineRecallManagementServiceResponse) SetBody(v *OnlineRecallManagementServiceResponseBody) *OnlineRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *OnlineRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
