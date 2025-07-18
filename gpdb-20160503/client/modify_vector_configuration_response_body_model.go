// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVectorConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyVectorConfigurationResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *ModifyVectorConfigurationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyVectorConfigurationResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ModifyVectorConfigurationResponseBody
	GetStatus() *bool
}

type ModifyVectorConfigurationResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message that is returned.
	//
	// This parameter is returned only if the request fails.
	//
	// example:
	//
	// Failed to modify vector configuration.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 107BE202-D1A2-479E-98E0-A8**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyVectorConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVectorConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVectorConfigurationResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyVectorConfigurationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyVectorConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVectorConfigurationResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ModifyVectorConfigurationResponseBody) SetDBInstanceId(v string) *ModifyVectorConfigurationResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) SetErrorMessage(v string) *ModifyVectorConfigurationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) SetRequestId(v string) *ModifyVectorConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) SetStatus(v bool) *ModifyVectorConfigurationResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyVectorConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
