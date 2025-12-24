// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertPlaybookReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertPlaybookReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertPlaybookReleaseResponse
	GetStatusCode() *int32
	SetBody(v *RevertPlaybookReleaseResponseBody) *RevertPlaybookReleaseResponse
	GetBody() *RevertPlaybookReleaseResponseBody
}

type RevertPlaybookReleaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertPlaybookReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertPlaybookReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertPlaybookReleaseResponse) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertPlaybookReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertPlaybookReleaseResponse) GetBody() *RevertPlaybookReleaseResponseBody {
	return s.Body
}

func (s *RevertPlaybookReleaseResponse) SetHeaders(v map[string]*string) *RevertPlaybookReleaseResponse {
	s.Headers = v
	return s
}

func (s *RevertPlaybookReleaseResponse) SetStatusCode(v int32) *RevertPlaybookReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertPlaybookReleaseResponse) SetBody(v *RevertPlaybookReleaseResponseBody) *RevertPlaybookReleaseResponse {
	s.Body = v
	return s
}

func (s *RevertPlaybookReleaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
