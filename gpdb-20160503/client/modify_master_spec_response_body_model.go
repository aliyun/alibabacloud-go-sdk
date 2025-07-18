// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMasterSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *ModifyMasterSpecResponseBody
	GetDbInstanceId() *string
	SetErrorMessage(v string) *ModifyMasterSpecResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyMasterSpecResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ModifyMasterSpecResponseBody
	GetStatus() *bool
}

type ModifyMasterSpecResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The returned message.
	//
	// This parameter is returned only if the request fails.
	//
	// example:
	//
	// ******
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMasterSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMasterSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMasterSpecResponseBody) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyMasterSpecResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyMasterSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMasterSpecResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ModifyMasterSpecResponseBody) SetDbInstanceId(v string) *ModifyMasterSpecResponseBody {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyMasterSpecResponseBody) SetErrorMessage(v string) *ModifyMasterSpecResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyMasterSpecResponseBody) SetRequestId(v string) *ModifyMasterSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMasterSpecResponseBody) SetStatus(v bool) *ModifyMasterSpecResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyMasterSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
