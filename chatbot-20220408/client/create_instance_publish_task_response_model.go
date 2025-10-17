// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancePublishTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstancePublishTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstancePublishTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstancePublishTaskResponseBody) *CreateInstancePublishTaskResponse
	GetBody() *CreateInstancePublishTaskResponseBody
}

type CreateInstancePublishTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstancePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstancePublishTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstancePublishTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstancePublishTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstancePublishTaskResponse) GetBody() *CreateInstancePublishTaskResponseBody {
	return s.Body
}

func (s *CreateInstancePublishTaskResponse) SetHeaders(v map[string]*string) *CreateInstancePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstancePublishTaskResponse) SetStatusCode(v int32) *CreateInstancePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstancePublishTaskResponse) SetBody(v *CreateInstancePublishTaskResponseBody) *CreateInstancePublishTaskResponse {
	s.Body = v
	return s
}

func (s *CreateInstancePublishTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
