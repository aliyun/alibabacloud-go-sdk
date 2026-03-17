// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWanSnatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagWanSnatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagWanSnatResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagWanSnatResponseBody) *ModifySagWanSnatResponse
	GetBody() *ModifySagWanSnatResponseBody
}

type ModifySagWanSnatResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagWanSnatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagWanSnatResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWanSnatResponse) GoString() string {
	return s.String()
}

func (s *ModifySagWanSnatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagWanSnatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagWanSnatResponse) GetBody() *ModifySagWanSnatResponseBody {
	return s.Body
}

func (s *ModifySagWanSnatResponse) SetHeaders(v map[string]*string) *ModifySagWanSnatResponse {
	s.Headers = v
	return s
}

func (s *ModifySagWanSnatResponse) SetStatusCode(v int32) *ModifySagWanSnatResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagWanSnatResponse) SetBody(v *ModifySagWanSnatResponseBody) *ModifySagWanSnatResponse {
	s.Body = v
	return s
}

func (s *ModifySagWanSnatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
