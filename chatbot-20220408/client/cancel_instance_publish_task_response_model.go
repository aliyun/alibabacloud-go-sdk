// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInstancePublishTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelInstancePublishTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelInstancePublishTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelInstancePublishTaskResponseBody) *CancelInstancePublishTaskResponse
	GetBody() *CancelInstancePublishTaskResponseBody
}

type CancelInstancePublishTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelInstancePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelInstancePublishTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelInstancePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelInstancePublishTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelInstancePublishTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelInstancePublishTaskResponse) GetBody() *CancelInstancePublishTaskResponseBody {
	return s.Body
}

func (s *CancelInstancePublishTaskResponse) SetHeaders(v map[string]*string) *CancelInstancePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelInstancePublishTaskResponse) SetStatusCode(v int32) *CancelInstancePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelInstancePublishTaskResponse) SetBody(v *CancelInstancePublishTaskResponseBody) *CancelInstancePublishTaskResponse {
	s.Body = v
	return s
}

func (s *CancelInstancePublishTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
