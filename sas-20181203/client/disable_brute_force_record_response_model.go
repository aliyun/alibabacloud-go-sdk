// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBruteForceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableBruteForceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableBruteForceRecordResponse
	GetStatusCode() *int32
	SetBody(v *DisableBruteForceRecordResponseBody) *DisableBruteForceRecordResponse
	GetBody() *DisableBruteForceRecordResponseBody
}

type DisableBruteForceRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableBruteForceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableBruteForceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableBruteForceRecordResponse) GoString() string {
	return s.String()
}

func (s *DisableBruteForceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableBruteForceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableBruteForceRecordResponse) GetBody() *DisableBruteForceRecordResponseBody {
	return s.Body
}

func (s *DisableBruteForceRecordResponse) SetHeaders(v map[string]*string) *DisableBruteForceRecordResponse {
	s.Headers = v
	return s
}

func (s *DisableBruteForceRecordResponse) SetStatusCode(v int32) *DisableBruteForceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableBruteForceRecordResponse) SetBody(v *DisableBruteForceRecordResponseBody) *DisableBruteForceRecordResponse {
	s.Body = v
	return s
}

func (s *DisableBruteForceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
