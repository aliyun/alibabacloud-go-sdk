// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPublishTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPublishTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPublishTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelPublishTaskResponseBody) *CancelPublishTaskResponse
	GetBody() *CancelPublishTaskResponseBody
}

type CancelPublishTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPublishTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPublishTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelPublishTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPublishTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPublishTaskResponse) GetBody() *CancelPublishTaskResponseBody {
	return s.Body
}

func (s *CancelPublishTaskResponse) SetHeaders(v map[string]*string) *CancelPublishTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelPublishTaskResponse) SetStatusCode(v int32) *CancelPublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPublishTaskResponse) SetBody(v *CancelPublishTaskResponseBody) *CancelPublishTaskResponse {
	s.Body = v
	return s
}

func (s *CancelPublishTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
