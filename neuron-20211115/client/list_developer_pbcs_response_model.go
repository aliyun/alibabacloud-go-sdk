// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeveloperPbcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeveloperPbcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeveloperPbcsResponse
	GetStatusCode() *int32
	SetBody(v *PbcListResult) *ListDeveloperPbcsResponse
	GetBody() *PbcListResult
}

type ListDeveloperPbcsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcListResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeveloperPbcsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeveloperPbcsResponse) GoString() string {
	return s.String()
}

func (s *ListDeveloperPbcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeveloperPbcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeveloperPbcsResponse) GetBody() *PbcListResult {
	return s.Body
}

func (s *ListDeveloperPbcsResponse) SetHeaders(v map[string]*string) *ListDeveloperPbcsResponse {
	s.Headers = v
	return s
}

func (s *ListDeveloperPbcsResponse) SetStatusCode(v int32) *ListDeveloperPbcsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeveloperPbcsResponse) SetBody(v *PbcListResult) *ListDeveloperPbcsResponse {
	s.Body = v
	return s
}

func (s *ListDeveloperPbcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
