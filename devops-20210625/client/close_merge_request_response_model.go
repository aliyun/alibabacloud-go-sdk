// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *CloseMergeRequestResponseBody) *CloseMergeRequestResponse
	GetBody() *CloseMergeRequestResponseBody
}

type CloseMergeRequestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *CloseMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseMergeRequestResponse) GetBody() *CloseMergeRequestResponseBody {
	return s.Body
}

func (s *CloseMergeRequestResponse) SetHeaders(v map[string]*string) *CloseMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *CloseMergeRequestResponse) SetStatusCode(v int32) *CloseMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseMergeRequestResponse) SetBody(v *CloseMergeRequestResponseBody) *CloseMergeRequestResponse {
	s.Body = v
	return s
}

func (s *CloseMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
