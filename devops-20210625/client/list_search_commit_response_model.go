// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchCommitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchCommitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchCommitResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchCommitResponseBody) *ListSearchCommitResponse
	GetBody() *ListSearchCommitResponseBody
}

type ListSearchCommitResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchCommitResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitResponse) GoString() string {
	return s.String()
}

func (s *ListSearchCommitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchCommitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchCommitResponse) GetBody() *ListSearchCommitResponseBody {
	return s.Body
}

func (s *ListSearchCommitResponse) SetHeaders(v map[string]*string) *ListSearchCommitResponse {
	s.Headers = v
	return s
}

func (s *ListSearchCommitResponse) SetStatusCode(v int32) *ListSearchCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchCommitResponse) SetBody(v *ListSearchCommitResponseBody) *ListSearchCommitResponse {
	s.Body = v
	return s
}

func (s *ListSearchCommitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
