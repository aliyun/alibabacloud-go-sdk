// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnsEipAddressAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEnsEipAddressAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEnsEipAddressAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEnsEipAddressAttributeResponseBody) *ModifyEnsEipAddressAttributeResponse
	GetBody() *ModifyEnsEipAddressAttributeResponseBody
}

type ModifyEnsEipAddressAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEnsEipAddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEnsEipAddressAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnsEipAddressAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyEnsEipAddressAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEnsEipAddressAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEnsEipAddressAttributeResponse) GetBody() *ModifyEnsEipAddressAttributeResponseBody {
	return s.Body
}

func (s *ModifyEnsEipAddressAttributeResponse) SetHeaders(v map[string]*string) *ModifyEnsEipAddressAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyEnsEipAddressAttributeResponse) SetStatusCode(v int32) *ModifyEnsEipAddressAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeResponse) SetBody(v *ModifyEnsEipAddressAttributeResponseBody) *ModifyEnsEipAddressAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyEnsEipAddressAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
