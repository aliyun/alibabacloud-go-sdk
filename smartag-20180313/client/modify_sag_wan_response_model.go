// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagWanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagWanResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagWanResponseBody) *ModifySagWanResponse
	GetBody() *ModifySagWanResponseBody
}

type ModifySagWanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagWanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagWanResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWanResponse) GoString() string {
	return s.String()
}

func (s *ModifySagWanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagWanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagWanResponse) GetBody() *ModifySagWanResponseBody {
	return s.Body
}

func (s *ModifySagWanResponse) SetHeaders(v map[string]*string) *ModifySagWanResponse {
	s.Headers = v
	return s
}

func (s *ModifySagWanResponse) SetStatusCode(v int32) *ModifySagWanResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagWanResponse) SetBody(v *ModifySagWanResponseBody) *ModifySagWanResponse {
	s.Body = v
	return s
}

func (s *ModifySagWanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
