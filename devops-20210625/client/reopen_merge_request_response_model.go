// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReopenMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReopenMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *ReopenMergeRequestResponseBody) *ReopenMergeRequestResponse
	GetBody() *ReopenMergeRequestResponseBody
}

type ReopenMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReopenMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReopenMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s ReopenMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *ReopenMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReopenMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReopenMergeRequestResponse) GetBody() *ReopenMergeRequestResponseBody {
	return s.Body
}

func (s *ReopenMergeRequestResponse) SetHeaders(v map[string]*string) *ReopenMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *ReopenMergeRequestResponse) SetStatusCode(v int32) *ReopenMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *ReopenMergeRequestResponse) SetBody(v *ReopenMergeRequestResponseBody) *ReopenMergeRequestResponse {
	s.Body = v
	return s
}

func (s *ReopenMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
