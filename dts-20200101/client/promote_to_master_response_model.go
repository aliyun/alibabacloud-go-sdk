// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteToMasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PromoteToMasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PromoteToMasterResponse
	GetStatusCode() *int32
	SetBody(v *PromoteToMasterResponseBody) *PromoteToMasterResponse
	GetBody() *PromoteToMasterResponseBody
}

type PromoteToMasterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PromoteToMasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PromoteToMasterResponse) String() string {
	return dara.Prettify(s)
}

func (s PromoteToMasterResponse) GoString() string {
	return s.String()
}

func (s *PromoteToMasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PromoteToMasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PromoteToMasterResponse) GetBody() *PromoteToMasterResponseBody {
	return s.Body
}

func (s *PromoteToMasterResponse) SetHeaders(v map[string]*string) *PromoteToMasterResponse {
	s.Headers = v
	return s
}

func (s *PromoteToMasterResponse) SetStatusCode(v int32) *PromoteToMasterResponse {
	s.StatusCode = &v
	return s
}

func (s *PromoteToMasterResponse) SetBody(v *PromoteToMasterResponseBody) *PromoteToMasterResponse {
	s.Body = v
	return s
}

func (s *PromoteToMasterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
