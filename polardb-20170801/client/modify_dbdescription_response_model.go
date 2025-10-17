// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBDescriptionResponseBody) *ModifyDBDescriptionResponse
	GetBody() *ModifyDBDescriptionResponseBody
}

type ModifyDBDescriptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBDescriptionResponse) GetBody() *ModifyDBDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDBDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBDescriptionResponse) SetStatusCode(v int32) *ModifyDBDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBDescriptionResponse) SetBody(v *ModifyDBDescriptionResponseBody) *ModifyDBDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
