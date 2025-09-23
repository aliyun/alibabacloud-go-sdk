// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEncryptionDBSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEncryptionDBSecretResponseBody
	GetDBClusterId() *string
	SetMessage(v string) *ModifyEncryptionDBSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyEncryptionDBSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyEncryptionDBSecretResponseBody
	GetSuccess() *bool
}

type ModifyEncryptionDBSecretResponseBody struct {
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4CE6DF97-AEA4-484F-906F-C407EE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyEncryptionDBSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEncryptionDBSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEncryptionDBSecretResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEncryptionDBSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyEncryptionDBSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEncryptionDBSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyEncryptionDBSecretResponseBody) SetDBClusterId(v string) *ModifyEncryptionDBSecretResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEncryptionDBSecretResponseBody) SetMessage(v string) *ModifyEncryptionDBSecretResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyEncryptionDBSecretResponseBody) SetRequestId(v string) *ModifyEncryptionDBSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEncryptionDBSecretResponseBody) SetSuccess(v bool) *ModifyEncryptionDBSecretResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyEncryptionDBSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
