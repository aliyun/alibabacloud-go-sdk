// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFuncSwitchRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFuncSwitchRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFuncSwitchRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFuncSwitchRecordResponseBody) *UpdateFuncSwitchRecordResponse
	GetBody() *UpdateFuncSwitchRecordResponseBody
}

type UpdateFuncSwitchRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFuncSwitchRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFuncSwitchRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFuncSwitchRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFuncSwitchRecordResponse) GetBody() *UpdateFuncSwitchRecordResponseBody {
	return s.Body
}

func (s *UpdateFuncSwitchRecordResponse) SetHeaders(v map[string]*string) *UpdateFuncSwitchRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateFuncSwitchRecordResponse) SetStatusCode(v int32) *UpdateFuncSwitchRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponse) SetBody(v *UpdateFuncSwitchRecordResponseBody) *UpdateFuncSwitchRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateFuncSwitchRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
