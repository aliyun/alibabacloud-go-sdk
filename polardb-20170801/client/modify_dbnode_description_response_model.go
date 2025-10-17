// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeDescriptionResponseBody) *ModifyDBNodeDescriptionResponse
	GetBody() *ModifyDBNodeDescriptionResponseBody
}

type ModifyDBNodeDescriptionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeDescriptionResponse) GetBody() *ModifyDBNodeDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDBNodeDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBNodeDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeDescriptionResponse) SetStatusCode(v int32) *ModifyDBNodeDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeDescriptionResponse) SetBody(v *ModifyDBNodeDescriptionResponseBody) *ModifyDBNodeDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
