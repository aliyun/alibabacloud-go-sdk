// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskWarningLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskWarningLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskWarningLineResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskWarningLineResponseBody) *ModifyDiskWarningLineResponse
	GetBody() *ModifyDiskWarningLineResponseBody
}

type ModifyDiskWarningLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskWarningLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskWarningLineResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskWarningLineResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskWarningLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskWarningLineResponse) GetBody() *ModifyDiskWarningLineResponseBody {
	return s.Body
}

func (s *ModifyDiskWarningLineResponse) SetHeaders(v map[string]*string) *ModifyDiskWarningLineResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskWarningLineResponse) SetStatusCode(v int32) *ModifyDiskWarningLineResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskWarningLineResponse) SetBody(v *ModifyDiskWarningLineResponseBody) *ModifyDiskWarningLineResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskWarningLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
