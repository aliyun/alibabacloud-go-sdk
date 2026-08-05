// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateOfflineTaskResponseBody) *CreateOfflineTaskResponse
	GetBody() *CreateOfflineTaskResponseBody
}

type CreateOfflineTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOfflineTaskResponse) GetBody() *CreateOfflineTaskResponseBody {
	return s.Body
}

func (s *CreateOfflineTaskResponse) SetHeaders(v map[string]*string) *CreateOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOfflineTaskResponse) SetStatusCode(v int32) *CreateOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOfflineTaskResponse) SetBody(v *CreateOfflineTaskResponseBody) *CreateOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *CreateOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
