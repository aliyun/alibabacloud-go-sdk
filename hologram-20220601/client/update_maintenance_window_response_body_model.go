// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintenanceWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateMaintenanceWindowResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateMaintenanceWindowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMaintenanceWindowResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateMaintenanceWindowResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateMaintenanceWindowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMaintenanceWindowResponseBody
	GetSuccess() *bool
}

type UpdateMaintenanceWindowResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMaintenanceWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintenanceWindowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaintenanceWindowResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateMaintenanceWindowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMaintenanceWindowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMaintenanceWindowResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateMaintenanceWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMaintenanceWindowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMaintenanceWindowResponseBody) SetData(v bool) *UpdateMaintenanceWindowResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMaintenanceWindowResponseBody) SetErrorCode(v string) *UpdateMaintenanceWindowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMaintenanceWindowResponseBody) SetErrorMessage(v string) *UpdateMaintenanceWindowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMaintenanceWindowResponseBody) SetHttpStatusCode(v string) *UpdateMaintenanceWindowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMaintenanceWindowResponseBody) SetRequestId(v string) *UpdateMaintenanceWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaintenanceWindowResponseBody) SetSuccess(v bool) *UpdateMaintenanceWindowResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMaintenanceWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
