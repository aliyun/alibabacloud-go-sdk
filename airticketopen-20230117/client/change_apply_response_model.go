// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeApplyResponse
	GetStatusCode() *int32
	SetBody(v *ChangeApplyResponseBody) *ChangeApplyResponse
	GetBody() *ChangeApplyResponseBody
}

type ChangeApplyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyResponse) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeApplyResponse) GetBody() *ChangeApplyResponseBody {
	return s.Body
}

func (s *ChangeApplyResponse) SetHeaders(v map[string]*string) *ChangeApplyResponse {
	s.Headers = v
	return s
}

func (s *ChangeApplyResponse) SetStatusCode(v int32) *ChangeApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeApplyResponse) SetBody(v *ChangeApplyResponseBody) *ChangeApplyResponse {
	s.Body = v
	return s
}

func (s *ChangeApplyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
