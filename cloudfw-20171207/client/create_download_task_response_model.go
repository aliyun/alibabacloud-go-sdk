// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDownloadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDownloadTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDownloadTaskResponseBody) *CreateDownloadTaskResponse
	GetBody() *CreateDownloadTaskResponseBody
}

type CreateDownloadTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDownloadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDownloadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDownloadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDownloadTaskResponse) GetBody() *CreateDownloadTaskResponseBody {
	return s.Body
}

func (s *CreateDownloadTaskResponse) SetHeaders(v map[string]*string) *CreateDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDownloadTaskResponse) SetStatusCode(v int32) *CreateDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDownloadTaskResponse) SetBody(v *CreateDownloadTaskResponseBody) *CreateDownloadTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDownloadTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
