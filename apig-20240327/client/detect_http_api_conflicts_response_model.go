// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectHttpApiConflictsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectHttpApiConflictsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectHttpApiConflictsResponse
	GetStatusCode() *int32
	SetBody(v *DetectHttpApiConflictsResponseBody) *DetectHttpApiConflictsResponse
	GetBody() *DetectHttpApiConflictsResponseBody
}

type DetectHttpApiConflictsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectHttpApiConflictsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectHttpApiConflictsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectHttpApiConflictsResponse) GoString() string {
	return s.String()
}

func (s *DetectHttpApiConflictsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectHttpApiConflictsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectHttpApiConflictsResponse) GetBody() *DetectHttpApiConflictsResponseBody {
	return s.Body
}

func (s *DetectHttpApiConflictsResponse) SetHeaders(v map[string]*string) *DetectHttpApiConflictsResponse {
	s.Headers = v
	return s
}

func (s *DetectHttpApiConflictsResponse) SetStatusCode(v int32) *DetectHttpApiConflictsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectHttpApiConflictsResponse) SetBody(v *DetectHttpApiConflictsResponseBody) *DetectHttpApiConflictsResponse {
	s.Body = v
	return s
}

func (s *DetectHttpApiConflictsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
