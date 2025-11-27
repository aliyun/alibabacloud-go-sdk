// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressPointCloudTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCompressPointCloudTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCompressPointCloudTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCompressPointCloudTaskResponseBody) *CreateCompressPointCloudTaskResponse
	GetBody() *CreateCompressPointCloudTaskResponseBody
}

type CreateCompressPointCloudTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCompressPointCloudTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCompressPointCloudTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressPointCloudTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCompressPointCloudTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCompressPointCloudTaskResponse) GetBody() *CreateCompressPointCloudTaskResponseBody {
	return s.Body
}

func (s *CreateCompressPointCloudTaskResponse) SetHeaders(v map[string]*string) *CreateCompressPointCloudTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCompressPointCloudTaskResponse) SetStatusCode(v int32) *CreateCompressPointCloudTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponse) SetBody(v *CreateCompressPointCloudTaskResponseBody) *CreateCompressPointCloudTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCompressPointCloudTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
