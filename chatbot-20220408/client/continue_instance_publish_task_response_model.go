// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueInstancePublishTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinueInstancePublishTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinueInstancePublishTaskResponse
	GetStatusCode() *int32
	SetBody(v *ContinueInstancePublishTaskResponseBody) *ContinueInstancePublishTaskResponse
	GetBody() *ContinueInstancePublishTaskResponseBody
}

type ContinueInstancePublishTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueInstancePublishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueInstancePublishTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinueInstancePublishTaskResponse) GoString() string {
	return s.String()
}

func (s *ContinueInstancePublishTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinueInstancePublishTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinueInstancePublishTaskResponse) GetBody() *ContinueInstancePublishTaskResponseBody {
	return s.Body
}

func (s *ContinueInstancePublishTaskResponse) SetHeaders(v map[string]*string) *ContinueInstancePublishTaskResponse {
	s.Headers = v
	return s
}

func (s *ContinueInstancePublishTaskResponse) SetStatusCode(v int32) *ContinueInstancePublishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueInstancePublishTaskResponse) SetBody(v *ContinueInstancePublishTaskResponseBody) *ContinueInstancePublishTaskResponse {
	s.Body = v
	return s
}

func (s *ContinueInstancePublishTaskResponse) Validate() error {
	return dara.Validate(s)
}
