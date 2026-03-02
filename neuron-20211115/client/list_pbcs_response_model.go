// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcsResponse
	GetStatusCode() *int32
	SetBody(v *PbcListResult) *ListPbcsResponse
	GetBody() *PbcListResult
}

type ListPbcsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcListResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcsResponse) GoString() string {
	return s.String()
}

func (s *ListPbcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcsResponse) GetBody() *PbcListResult {
	return s.Body
}

func (s *ListPbcsResponse) SetHeaders(v map[string]*string) *ListPbcsResponse {
	s.Headers = v
	return s
}

func (s *ListPbcsResponse) SetStatusCode(v int32) *ListPbcsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcsResponse) SetBody(v *PbcListResult) *ListPbcsResponse {
	s.Body = v
	return s
}

func (s *ListPbcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
