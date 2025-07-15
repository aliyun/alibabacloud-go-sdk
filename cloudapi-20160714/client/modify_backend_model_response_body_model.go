// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackendModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *ModifyBackendModelResponseBody
	GetOperationId() *string
	SetRequestId(v string) *ModifyBackendModelResponseBody
	GetRequestId() *string
}

type ModifyBackendModelResponseBody struct {
	// example:
	//
	// c16a1880f5164d779f6a54f64d997cd9
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// example:
	//
	// 19B89B04-418B-55EE-94A8-6B42CA06002A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackendModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackendModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackendModelResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *ModifyBackendModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackendModelResponseBody) SetOperationId(v string) *ModifyBackendModelResponseBody {
	s.OperationId = &v
	return s
}

func (s *ModifyBackendModelResponseBody) SetRequestId(v string) *ModifyBackendModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackendModelResponseBody) Validate() error {
	return dara.Validate(s)
}
