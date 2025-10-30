// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMasterSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMasterSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMasterSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMasterSpecResponseBody) *ModifyMasterSpecResponse
	GetBody() *ModifyMasterSpecResponseBody
}

type ModifyMasterSpecResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMasterSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMasterSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMasterSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyMasterSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMasterSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMasterSpecResponse) GetBody() *ModifyMasterSpecResponseBody {
	return s.Body
}

func (s *ModifyMasterSpecResponse) SetHeaders(v map[string]*string) *ModifyMasterSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyMasterSpecResponse) SetStatusCode(v int32) *ModifyMasterSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMasterSpecResponse) SetBody(v *ModifyMasterSpecResponseBody) *ModifyMasterSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyMasterSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
