// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppInstanceSpecResponseBody) *ModifyAppInstanceSpecResponse
	GetBody() *ModifyAppInstanceSpecResponseBody
}

type ModifyAppInstanceSpecResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppInstanceSpecResponse) GetBody() *ModifyAppInstanceSpecResponseBody {
	return s.Body
}

func (s *ModifyAppInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceSpecResponse) SetStatusCode(v int32) *ModifyAppInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceSpecResponse) SetBody(v *ModifyAppInstanceSpecResponseBody) *ModifyAppInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyAppInstanceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
