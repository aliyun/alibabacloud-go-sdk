// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCDiskSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCDiskSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCDiskSpecResponseBody) *ModifyRCDiskSpecResponse
	GetBody() *ModifyRCDiskSpecResponseBody
}

type ModifyRCDiskSpecResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCDiskSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCDiskSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCDiskSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCDiskSpecResponse) GetBody() *ModifyRCDiskSpecResponseBody {
	return s.Body
}

func (s *ModifyRCDiskSpecResponse) SetHeaders(v map[string]*string) *ModifyRCDiskSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCDiskSpecResponse) SetStatusCode(v int32) *ModifyRCDiskSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCDiskSpecResponse) SetBody(v *ModifyRCDiskSpecResponseBody) *ModifyRCDiskSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyRCDiskSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
