// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeTwoWayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *UpgradeTwoWayResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpgradeTwoWayResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *UpgradeTwoWayResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpgradeTwoWayResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpgradeTwoWayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpgradeTwoWayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeTwoWayResponseBody
	GetSuccess() *bool
}

type UpgradeTwoWayResponseBody struct {
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and that of the **DynamicMessage*	- parameter is **InstanceId**, the specified **InstanceId*	- parameter is invalid.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2D3B4615-923F-49AA-AF21-6D8E3967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**: The call was successful.
	//
	// 	- **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeTwoWayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeTwoWayResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeTwoWayResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpgradeTwoWayResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpgradeTwoWayResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpgradeTwoWayResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpgradeTwoWayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpgradeTwoWayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeTwoWayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeTwoWayResponseBody) SetDynamicCode(v string) *UpgradeTwoWayResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetDynamicMessage(v string) *UpgradeTwoWayResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetErrCode(v string) *UpgradeTwoWayResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetErrMessage(v string) *UpgradeTwoWayResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetHttpStatusCode(v int32) *UpgradeTwoWayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetRequestId(v string) *UpgradeTwoWayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetSuccess(v bool) *UpgradeTwoWayResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) Validate() error {
	return dara.Validate(s)
}
