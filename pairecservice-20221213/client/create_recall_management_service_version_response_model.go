// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecallManagementServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecallManagementServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecallManagementServiceVersionResponseBody) *CreateRecallManagementServiceVersionResponse
	GetBody() *CreateRecallManagementServiceVersionResponseBody
}

type CreateRecallManagementServiceVersionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecallManagementServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecallManagementServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecallManagementServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecallManagementServiceVersionResponse) GetBody() *CreateRecallManagementServiceVersionResponseBody {
	return s.Body
}

func (s *CreateRecallManagementServiceVersionResponse) SetHeaders(v map[string]*string) *CreateRecallManagementServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateRecallManagementServiceVersionResponse) SetStatusCode(v int32) *CreateRecallManagementServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecallManagementServiceVersionResponse) SetBody(v *CreateRecallManagementServiceVersionResponseBody) *CreateRecallManagementServiceVersionResponse {
	s.Body = v
	return s
}

func (s *CreateRecallManagementServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
