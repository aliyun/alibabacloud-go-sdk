// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterCnnfStatusUserConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterCnnfStatusUserConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterCnnfStatusUserConfirmResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterCnnfStatusUserConfirmResponseBody) *ModifyClusterCnnfStatusUserConfirmResponse
	GetBody() *ModifyClusterCnnfStatusUserConfirmResponseBody
}

type ModifyClusterCnnfStatusUserConfirmResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterCnnfStatusUserConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterCnnfStatusUserConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterCnnfStatusUserConfirmResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) GetBody() *ModifyClusterCnnfStatusUserConfirmResponseBody {
	return s.Body
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) SetHeaders(v map[string]*string) *ModifyClusterCnnfStatusUserConfirmResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) SetStatusCode(v int32) *ModifyClusterCnnfStatusUserConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) SetBody(v *ModifyClusterCnnfStatusUserConfirmResponseBody) *ModifyClusterCnnfStatusUserConfirmResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterCnnfStatusUserConfirmResponse) Validate() error {
	return dara.Validate(s)
}
