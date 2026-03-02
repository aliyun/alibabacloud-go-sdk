// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvokePbcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInvokePbcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInvokePbcsResponse
	GetStatusCode() *int32
	SetBody(v *PbcListResult) *ListInvokePbcsResponse
	GetBody() *PbcListResult
}

type ListInvokePbcsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcListResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInvokePbcsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInvokePbcsResponse) GoString() string {
	return s.String()
}

func (s *ListInvokePbcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInvokePbcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInvokePbcsResponse) GetBody() *PbcListResult {
	return s.Body
}

func (s *ListInvokePbcsResponse) SetHeaders(v map[string]*string) *ListInvokePbcsResponse {
	s.Headers = v
	return s
}

func (s *ListInvokePbcsResponse) SetStatusCode(v int32) *ListInvokePbcsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInvokePbcsResponse) SetBody(v *PbcListResult) *ListInvokePbcsResponse {
	s.Body = v
	return s
}

func (s *ListInvokePbcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
