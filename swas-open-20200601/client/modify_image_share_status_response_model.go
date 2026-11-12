// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageShareStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImageShareStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImageShareStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImageShareStatusResponseBody) *ModifyImageShareStatusResponse
	GetBody() *ModifyImageShareStatusResponseBody
}

type ModifyImageShareStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageShareStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageShareStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImageShareStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImageShareStatusResponse) GetBody() *ModifyImageShareStatusResponseBody {
	return s.Body
}

func (s *ModifyImageShareStatusResponse) SetHeaders(v map[string]*string) *ModifyImageShareStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageShareStatusResponse) SetStatusCode(v int32) *ModifyImageShareStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageShareStatusResponse) SetBody(v *ModifyImageShareStatusResponseBody) *ModifyImageShareStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyImageShareStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
