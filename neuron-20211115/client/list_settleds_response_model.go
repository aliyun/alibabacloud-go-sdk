// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSettledsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSettledsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSettledsResponse
	GetStatusCode() *int32
	SetBody(v *SettledPageResult) *ListSettledsResponse
	GetBody() *SettledPageResult
}

type ListSettledsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SettledPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSettledsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSettledsResponse) GoString() string {
	return s.String()
}

func (s *ListSettledsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSettledsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSettledsResponse) GetBody() *SettledPageResult {
	return s.Body
}

func (s *ListSettledsResponse) SetHeaders(v map[string]*string) *ListSettledsResponse {
	s.Headers = v
	return s
}

func (s *ListSettledsResponse) SetStatusCode(v int32) *ListSettledsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSettledsResponse) SetBody(v *SettledPageResult) *ListSettledsResponse {
	s.Body = v
	return s
}

func (s *ListSettledsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
