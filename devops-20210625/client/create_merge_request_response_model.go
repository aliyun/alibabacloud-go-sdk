// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateMergeRequestResponseBody) *CreateMergeRequestResponse
	GetBody() *CreateMergeRequestResponseBody
}

type CreateMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMergeRequestResponse) GetBody() *CreateMergeRequestResponseBody {
	return s.Body
}

func (s *CreateMergeRequestResponse) SetHeaders(v map[string]*string) *CreateMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateMergeRequestResponse) SetStatusCode(v int32) *CreateMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMergeRequestResponse) SetBody(v *CreateMergeRequestResponseBody) *CreateMergeRequestResponse {
	s.Body = v
	return s
}

func (s *CreateMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
