// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *ModifyDBInstanceConfigResponseBody
	GetDbInstanceId() *string
	SetErrorMessage(v string) *ModifyDBInstanceConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyDBInstanceConfigResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ModifyDBInstanceConfigResponseBody
	GetStatus() *bool
}

type ModifyDBInstanceConfigResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The error message returned if the operation fails.
	//
	// example:
	//
	// \\"error message\\"
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34B32A0A-08EF-4A87-B6BE-CDD9F56FC3AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation. Valid values:
	//
	// 	- **0**: The operation failed.
	//
	// 	- **1**: The operation is successful.
	//
	// example:
	//
	// 1
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDBInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponseBody) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyDBInstanceConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyDBInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceConfigResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ModifyDBInstanceConfigResponseBody) SetDbInstanceId(v string) *ModifyDBInstanceConfigResponseBody {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetErrorMessage(v string) *ModifyDBInstanceConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetRequestId(v string) *ModifyDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetStatus(v bool) *ModifyDBInstanceConfigResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
