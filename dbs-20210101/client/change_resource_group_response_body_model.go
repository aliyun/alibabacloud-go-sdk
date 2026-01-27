// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeResourceGroupResponseBody
	GetCode() *string
	SetData(v string) *ChangeResourceGroupResponseBody
	GetData() *string
	SetErrCode(v string) *ChangeResourceGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChangeResourceGroupResponseBody
	GetErrMessage() *string
	SetMessage(v string) *ChangeResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ChangeResourceGroupResponseBody
	GetSuccess() *string
}

type ChangeResourceGroupResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the resource was successfully moved. Valid values:
	//
	// 	- **true**: The resource was successfully moved.
	//
	// 	- **false**: The resource failed to be moved.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Request.Forbidden
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// RAM DENY
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The additional information.
	//
	// example:
	//
	// The resource group is forbidden to operate
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04EBD9F5-F06F-5302-8499-005C72*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeResourceGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *ChangeResourceGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChangeResourceGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChangeResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v string) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetErrCode(v string) *ChangeResourceGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetErrMessage(v string) *ChangeResourceGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v string) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
