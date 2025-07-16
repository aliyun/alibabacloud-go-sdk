// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyActiveOperationMaintainConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyActiveOperationMaintainConfResponse
	GetStatusCode() *int32
	SetBody(v *ModifyActiveOperationMaintainConfResponseBody) *ModifyActiveOperationMaintainConfResponse
	GetBody() *ModifyActiveOperationMaintainConfResponseBody
}

type ModifyActiveOperationMaintainConfResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationMaintainConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationMaintainConfResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyActiveOperationMaintainConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyActiveOperationMaintainConfResponse) GetBody() *ModifyActiveOperationMaintainConfResponseBody {
	return s.Body
}

func (s *ModifyActiveOperationMaintainConfResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationMaintainConfResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationMaintainConfResponse) SetStatusCode(v int32) *ModifyActiveOperationMaintainConfResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfResponse) SetBody(v *ModifyActiveOperationMaintainConfResponseBody) *ModifyActiveOperationMaintainConfResponse {
	s.Body = v
	return s
}

func (s *ModifyActiveOperationMaintainConfResponse) Validate() error {
	return dara.Validate(s)
}
