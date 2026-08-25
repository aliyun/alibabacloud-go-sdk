// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoftwarelibDistributeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSoftwarelibDistributeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSoftwarelibDistributeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSoftwarelibDistributeTaskResponseBody) *CreateSoftwarelibDistributeTaskResponse
	GetBody() *CreateSoftwarelibDistributeTaskResponseBody
}

type CreateSoftwarelibDistributeTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSoftwarelibDistributeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSoftwarelibDistributeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibDistributeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibDistributeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSoftwarelibDistributeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSoftwarelibDistributeTaskResponse) GetBody() *CreateSoftwarelibDistributeTaskResponseBody {
	return s.Body
}

func (s *CreateSoftwarelibDistributeTaskResponse) SetHeaders(v map[string]*string) *CreateSoftwarelibDistributeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponse) SetStatusCode(v int32) *CreateSoftwarelibDistributeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponse) SetBody(v *CreateSoftwarelibDistributeTaskResponseBody) *CreateSoftwarelibDistributeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
