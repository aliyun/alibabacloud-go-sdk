// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDetectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageDetectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageDetectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageDetectionTaskResponseBody) *CreateImageDetectionTaskResponse
	GetBody() *CreateImageDetectionTaskResponseBody
}

type CreateImageDetectionTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageDetectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageDetectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageDetectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageDetectionTaskResponse) GetBody() *CreateImageDetectionTaskResponseBody {
	return s.Body
}

func (s *CreateImageDetectionTaskResponse) SetHeaders(v map[string]*string) *CreateImageDetectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageDetectionTaskResponse) SetStatusCode(v int32) *CreateImageDetectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageDetectionTaskResponse) SetBody(v *CreateImageDetectionTaskResponseBody) *CreateImageDetectionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageDetectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
