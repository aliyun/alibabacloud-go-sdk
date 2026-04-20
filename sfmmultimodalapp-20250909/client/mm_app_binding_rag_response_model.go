// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingRagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MmAppBindingRagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MmAppBindingRagResponse
	GetStatusCode() *int32
	SetBody(v *MmAppBindingRagResponseBody) *MmAppBindingRagResponse
	GetBody() *MmAppBindingRagResponseBody
}

type MmAppBindingRagResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MmAppBindingRagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MmAppBindingRagResponse) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingRagResponse) GoString() string {
	return s.String()
}

func (s *MmAppBindingRagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MmAppBindingRagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MmAppBindingRagResponse) GetBody() *MmAppBindingRagResponseBody {
	return s.Body
}

func (s *MmAppBindingRagResponse) SetHeaders(v map[string]*string) *MmAppBindingRagResponse {
	s.Headers = v
	return s
}

func (s *MmAppBindingRagResponse) SetStatusCode(v int32) *MmAppBindingRagResponse {
	s.StatusCode = &v
	return s
}

func (s *MmAppBindingRagResponse) SetBody(v *MmAppBindingRagResponseBody) *MmAppBindingRagResponse {
	s.Body = v
	return s
}

func (s *MmAppBindingRagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
