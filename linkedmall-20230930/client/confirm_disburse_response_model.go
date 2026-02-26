// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDisburseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmDisburseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmDisburseResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmDisburseResult) *ConfirmDisburseResponse
	GetBody() *ConfirmDisburseResult
}

type ConfirmDisburseResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmDisburseResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDisburseResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDisburseResponse) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmDisburseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmDisburseResponse) GetBody() *ConfirmDisburseResult {
	return s.Body
}

func (s *ConfirmDisburseResponse) SetHeaders(v map[string]*string) *ConfirmDisburseResponse {
	s.Headers = v
	return s
}

func (s *ConfirmDisburseResponse) SetStatusCode(v int32) *ConfirmDisburseResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmDisburseResponse) SetBody(v *ConfirmDisburseResult) *ConfirmDisburseResponse {
	s.Body = v
	return s
}

func (s *ConfirmDisburseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
