// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecallManagementServiceResponseBody) *DeleteRecallManagementServiceResponse
	GetBody() *DeleteRecallManagementServiceResponseBody
}

type DeleteRecallManagementServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecallManagementServiceResponse) GetBody() *DeleteRecallManagementServiceResponseBody {
	return s.Body
}

func (s *DeleteRecallManagementServiceResponse) SetHeaders(v map[string]*string) *DeleteRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecallManagementServiceResponse) SetStatusCode(v int32) *DeleteRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecallManagementServiceResponse) SetBody(v *DeleteRecallManagementServiceResponseBody) *DeleteRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
