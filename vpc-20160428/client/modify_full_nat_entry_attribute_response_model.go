// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullNatEntryAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFullNatEntryAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFullNatEntryAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFullNatEntryAttributeResponseBody) *ModifyFullNatEntryAttributeResponse
	GetBody() *ModifyFullNatEntryAttributeResponseBody
}

type ModifyFullNatEntryAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFullNatEntryAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFullNatEntryAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullNatEntryAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyFullNatEntryAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFullNatEntryAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFullNatEntryAttributeResponse) GetBody() *ModifyFullNatEntryAttributeResponseBody {
	return s.Body
}

func (s *ModifyFullNatEntryAttributeResponse) SetHeaders(v map[string]*string) *ModifyFullNatEntryAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyFullNatEntryAttributeResponse) SetStatusCode(v int32) *ModifyFullNatEntryAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFullNatEntryAttributeResponse) SetBody(v *ModifyFullNatEntryAttributeResponseBody) *ModifyFullNatEntryAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyFullNatEntryAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
