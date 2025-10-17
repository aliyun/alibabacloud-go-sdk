// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePublishTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePublishTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatePublishTaskResponseBody) *CreatePublishTaskResponse
	GetBody() *CreatePublishTaskResponseBody
}

type CreatePublishTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublishTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePublishTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePublishTaskResponse) GetBody() *CreatePublishTaskResponseBody {
	return s.Body
}

func (s *CreatePublishTaskResponse) SetHeaders(v map[string]*string) *CreatePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePublishTaskResponse) SetStatusCode(v int32) *CreatePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublishTaskResponse) SetBody(v *CreatePublishTaskResponseBody) *CreatePublishTaskResponse {
	s.Body = v
	return s
}

func (s *CreatePublishTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
