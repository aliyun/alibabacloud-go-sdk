// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecallManagementTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecallManagementTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecallManagementTableResponseBody) *UpdateRecallManagementTableResponse
	GetBody() *UpdateRecallManagementTableResponseBody
}

type UpdateRecallManagementTableResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecallManagementTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecallManagementTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecallManagementTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecallManagementTableResponse) GetBody() *UpdateRecallManagementTableResponseBody {
	return s.Body
}

func (s *UpdateRecallManagementTableResponse) SetHeaders(v map[string]*string) *UpdateRecallManagementTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecallManagementTableResponse) SetStatusCode(v int32) *UpdateRecallManagementTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecallManagementTableResponse) SetBody(v *UpdateRecallManagementTableResponseBody) *UpdateRecallManagementTableResponse {
	s.Body = v
	return s
}

func (s *UpdateRecallManagementTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
