// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecallManagementServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecallManagementServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecallManagementServiceVersionResponseBody) *DeleteRecallManagementServiceVersionResponse
	GetBody() *DeleteRecallManagementServiceVersionResponseBody
}

type DeleteRecallManagementServiceVersionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecallManagementServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecallManagementServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecallManagementServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecallManagementServiceVersionResponse) GetBody() *DeleteRecallManagementServiceVersionResponseBody {
	return s.Body
}

func (s *DeleteRecallManagementServiceVersionResponse) SetHeaders(v map[string]*string) *DeleteRecallManagementServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecallManagementServiceVersionResponse) SetStatusCode(v int32) *DeleteRecallManagementServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecallManagementServiceVersionResponse) SetBody(v *DeleteRecallManagementServiceVersionResponseBody) *DeleteRecallManagementServiceVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteRecallManagementServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
