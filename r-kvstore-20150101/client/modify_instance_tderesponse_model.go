// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTDEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceTDEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceTDEResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceTDEResponseBody) *ModifyInstanceTDEResponse
	GetBody() *ModifyInstanceTDEResponseBody
}

type ModifyInstanceTDEResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceTDEResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTDEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceTDEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceTDEResponse) GetBody() *ModifyInstanceTDEResponseBody {
	return s.Body
}

func (s *ModifyInstanceTDEResponse) SetHeaders(v map[string]*string) *ModifyInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceTDEResponse) SetStatusCode(v int32) *ModifyInstanceTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceTDEResponse) SetBody(v *ModifyInstanceTDEResponseBody) *ModifyInstanceTDEResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceTDEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
