// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineDatabaseTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOnlineDatabaseTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOnlineDatabaseTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateOnlineDatabaseTaskResponseBody) *CreateOnlineDatabaseTaskResponse
	GetBody() *CreateOnlineDatabaseTaskResponseBody
}

type CreateOnlineDatabaseTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOnlineDatabaseTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOnlineDatabaseTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineDatabaseTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOnlineDatabaseTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOnlineDatabaseTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOnlineDatabaseTaskResponse) GetBody() *CreateOnlineDatabaseTaskResponseBody {
	return s.Body
}

func (s *CreateOnlineDatabaseTaskResponse) SetHeaders(v map[string]*string) *CreateOnlineDatabaseTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOnlineDatabaseTaskResponse) SetStatusCode(v int32) *CreateOnlineDatabaseTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOnlineDatabaseTaskResponse) SetBody(v *CreateOnlineDatabaseTaskResponseBody) *CreateOnlineDatabaseTaskResponse {
	s.Body = v
	return s
}

func (s *CreateOnlineDatabaseTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
