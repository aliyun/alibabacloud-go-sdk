// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShutdownPolicyRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetShutdownPolicyRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetShutdownPolicyRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetShutdownPolicyRecordResponseBody) *GetShutdownPolicyRecordResponse
	GetBody() *GetShutdownPolicyRecordResponseBody
}

type GetShutdownPolicyRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetShutdownPolicyRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShutdownPolicyRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetShutdownPolicyRecordResponse) GoString() string {
	return s.String()
}

func (s *GetShutdownPolicyRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetShutdownPolicyRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetShutdownPolicyRecordResponse) GetBody() *GetShutdownPolicyRecordResponseBody {
	return s.Body
}

func (s *GetShutdownPolicyRecordResponse) SetHeaders(v map[string]*string) *GetShutdownPolicyRecordResponse {
	s.Headers = v
	return s
}

func (s *GetShutdownPolicyRecordResponse) SetStatusCode(v int32) *GetShutdownPolicyRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShutdownPolicyRecordResponse) SetBody(v *GetShutdownPolicyRecordResponseBody) *GetShutdownPolicyRecordResponse {
	s.Body = v
	return s
}

func (s *GetShutdownPolicyRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
