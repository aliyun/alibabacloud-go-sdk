// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndividuationTextTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIndividuationTextTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIndividuationTextTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateIndividuationTextTaskResponseBody) *CreateIndividuationTextTaskResponse
	GetBody() *CreateIndividuationTextTaskResponseBody
}

type CreateIndividuationTextTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndividuationTextTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndividuationTextTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIndividuationTextTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateIndividuationTextTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIndividuationTextTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIndividuationTextTaskResponse) GetBody() *CreateIndividuationTextTaskResponseBody {
	return s.Body
}

func (s *CreateIndividuationTextTaskResponse) SetHeaders(v map[string]*string) *CreateIndividuationTextTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateIndividuationTextTaskResponse) SetStatusCode(v int32) *CreateIndividuationTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndividuationTextTaskResponse) SetBody(v *CreateIndividuationTextTaskResponseBody) *CreateIndividuationTextTaskResponse {
	s.Body = v
	return s
}

func (s *CreateIndividuationTextTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
