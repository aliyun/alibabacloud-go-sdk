// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceVersionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecallManagementServiceVersionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecallManagementServiceVersionConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecallManagementServiceVersionConfigResponseBody) *DeleteRecallManagementServiceVersionConfigResponse
	GetBody() *DeleteRecallManagementServiceVersionConfigResponseBody
}

type DeleteRecallManagementServiceVersionConfigResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecallManagementServiceVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecallManagementServiceVersionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) GetBody() *DeleteRecallManagementServiceVersionConfigResponseBody {
	return s.Body
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) SetHeaders(v map[string]*string) *DeleteRecallManagementServiceVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) SetStatusCode(v int32) *DeleteRecallManagementServiceVersionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) SetBody(v *DeleteRecallManagementServiceVersionConfigResponseBody) *DeleteRecallManagementServiceVersionConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteRecallManagementServiceVersionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
