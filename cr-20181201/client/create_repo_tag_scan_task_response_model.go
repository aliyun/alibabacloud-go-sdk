// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTagScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoTagScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoTagScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoTagScanTaskResponseBody) *CreateRepoTagScanTaskResponse
	GetBody() *CreateRepoTagScanTaskResponseBody
}

type CreateRepoTagScanTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoTagScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoTagScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTagScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoTagScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoTagScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoTagScanTaskResponse) GetBody() *CreateRepoTagScanTaskResponseBody {
	return s.Body
}

func (s *CreateRepoTagScanTaskResponse) SetHeaders(v map[string]*string) *CreateRepoTagScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoTagScanTaskResponse) SetStatusCode(v int32) *CreateRepoTagScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoTagScanTaskResponse) SetBody(v *CreateRepoTagScanTaskResponseBody) *CreateRepoTagScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRepoTagScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
