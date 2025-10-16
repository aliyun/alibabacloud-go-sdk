// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodePoolAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodePoolAmountResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodePoolAmountResponseBody) *ModifyNodePoolAmountResponse
	GetBody() *ModifyNodePoolAmountResponseBody
}

type ModifyNodePoolAmountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodePoolAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodePoolAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAmountResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodePoolAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodePoolAmountResponse) GetBody() *ModifyNodePoolAmountResponseBody {
	return s.Body
}

func (s *ModifyNodePoolAmountResponse) SetHeaders(v map[string]*string) *ModifyNodePoolAmountResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolAmountResponse) SetStatusCode(v int32) *ModifyNodePoolAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolAmountResponse) SetBody(v *ModifyNodePoolAmountResponseBody) *ModifyNodePoolAmountResponse {
	s.Body = v
	return s
}

func (s *ModifyNodePoolAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
