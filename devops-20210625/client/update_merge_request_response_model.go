// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMergeRequestResponseBody) *UpdateMergeRequestResponse
	GetBody() *UpdateMergeRequestResponseBody
}

type UpdateMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMergeRequestResponse) GetBody() *UpdateMergeRequestResponseBody {
	return s.Body
}

func (s *UpdateMergeRequestResponse) SetHeaders(v map[string]*string) *UpdateMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *UpdateMergeRequestResponse) SetStatusCode(v int32) *UpdateMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMergeRequestResponse) SetBody(v *UpdateMergeRequestResponseBody) *UpdateMergeRequestResponse {
	s.Body = v
	return s
}

func (s *UpdateMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
