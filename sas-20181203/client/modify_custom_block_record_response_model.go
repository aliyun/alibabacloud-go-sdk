// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomBlockRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomBlockRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomBlockRecordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomBlockRecordResponseBody) *ModifyCustomBlockRecordResponse
	GetBody() *ModifyCustomBlockRecordResponseBody
}

type ModifyCustomBlockRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomBlockRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomBlockRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomBlockRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomBlockRecordResponse) GetBody() *ModifyCustomBlockRecordResponseBody {
	return s.Body
}

func (s *ModifyCustomBlockRecordResponse) SetHeaders(v map[string]*string) *ModifyCustomBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomBlockRecordResponse) SetStatusCode(v int32) *ModifyCustomBlockRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomBlockRecordResponse) SetBody(v *ModifyCustomBlockRecordResponseBody) *ModifyCustomBlockRecordResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomBlockRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
