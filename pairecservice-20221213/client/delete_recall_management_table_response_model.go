// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecallManagementTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecallManagementTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecallManagementTableResponseBody) *DeleteRecallManagementTableResponse
	GetBody() *DeleteRecallManagementTableResponseBody
}

type DeleteRecallManagementTableResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecallManagementTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecallManagementTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecallManagementTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecallManagementTableResponse) GetBody() *DeleteRecallManagementTableResponseBody {
	return s.Body
}

func (s *DeleteRecallManagementTableResponse) SetHeaders(v map[string]*string) *DeleteRecallManagementTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecallManagementTableResponse) SetStatusCode(v int32) *DeleteRecallManagementTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecallManagementTableResponse) SetBody(v *DeleteRecallManagementTableResponseBody) *DeleteRecallManagementTableResponse {
	s.Body = v
	return s
}

func (s *DeleteRecallManagementTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
