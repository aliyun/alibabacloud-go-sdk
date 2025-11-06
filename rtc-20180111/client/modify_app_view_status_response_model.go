// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppViewStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppViewStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppViewStatusResponseBody) *ModifyAppViewStatusResponse
	GetBody() *ModifyAppViewStatusResponseBody
}

type ModifyAppViewStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppViewStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppViewStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppViewStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppViewStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppViewStatusResponse) GetBody() *ModifyAppViewStatusResponseBody {
	return s.Body
}

func (s *ModifyAppViewStatusResponse) SetHeaders(v map[string]*string) *ModifyAppViewStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppViewStatusResponse) SetStatusCode(v int32) *ModifyAppViewStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppViewStatusResponse) SetBody(v *ModifyAppViewStatusResponseBody) *ModifyAppViewStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyAppViewStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
