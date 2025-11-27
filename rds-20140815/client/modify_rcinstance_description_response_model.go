// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceDescriptionResponseBody) *ModifyRCInstanceDescriptionResponse
	GetBody() *ModifyRCInstanceDescriptionResponseBody
}

type ModifyRCInstanceDescriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceDescriptionResponse) GetBody() *ModifyRCInstanceDescriptionResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyRCInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceDescriptionResponse) SetBody(v *ModifyRCInstanceDescriptionResponseBody) *ModifyRCInstanceDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
