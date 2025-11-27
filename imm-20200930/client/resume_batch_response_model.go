// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeBatchResponse
	GetStatusCode() *int32
	SetBody(v *ResumeBatchResponseBody) *ResumeBatchResponse
	GetBody() *ResumeBatchResponseBody
}

type ResumeBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeBatchResponse) GoString() string {
	return s.String()
}

func (s *ResumeBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeBatchResponse) GetBody() *ResumeBatchResponseBody {
	return s.Body
}

func (s *ResumeBatchResponse) SetHeaders(v map[string]*string) *ResumeBatchResponse {
	s.Headers = v
	return s
}

func (s *ResumeBatchResponse) SetStatusCode(v int32) *ResumeBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeBatchResponse) SetBody(v *ResumeBatchResponseBody) *ResumeBatchResponse {
	s.Body = v
	return s
}

func (s *ResumeBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
