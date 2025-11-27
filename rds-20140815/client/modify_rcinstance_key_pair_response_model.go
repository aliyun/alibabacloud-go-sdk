// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceKeyPairResponseBody) *ModifyRCInstanceKeyPairResponse
	GetBody() *ModifyRCInstanceKeyPairResponseBody
}

type ModifyRCInstanceKeyPairResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceKeyPairResponse) GetBody() *ModifyRCInstanceKeyPairResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceKeyPairResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceKeyPairResponse) SetStatusCode(v int32) *ModifyRCInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceKeyPairResponse) SetBody(v *ModifyRCInstanceKeyPairResponseBody) *ModifyRCInstanceKeyPairResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
