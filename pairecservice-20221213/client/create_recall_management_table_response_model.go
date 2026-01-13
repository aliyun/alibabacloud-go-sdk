// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecallManagementTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecallManagementTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecallManagementTableResponseBody) *CreateRecallManagementTableResponse
	GetBody() *CreateRecallManagementTableResponseBody
}

type CreateRecallManagementTableResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecallManagementTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecallManagementTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementTableResponse) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecallManagementTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecallManagementTableResponse) GetBody() *CreateRecallManagementTableResponseBody {
	return s.Body
}

func (s *CreateRecallManagementTableResponse) SetHeaders(v map[string]*string) *CreateRecallManagementTableResponse {
	s.Headers = v
	return s
}

func (s *CreateRecallManagementTableResponse) SetStatusCode(v int32) *CreateRecallManagementTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecallManagementTableResponse) SetBody(v *CreateRecallManagementTableResponseBody) *CreateRecallManagementTableResponse {
	s.Body = v
	return s
}

func (s *CreateRecallManagementTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
