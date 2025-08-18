// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAppInstanceModels(v []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels) *DeleteAppInstancesResponseBody
	GetDeleteAppInstanceModels() []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels
	SetRequestId(v string) *DeleteAppInstancesResponseBody
	GetRequestId() *string
}

type DeleteAppInstancesResponseBody struct {
	// The data returned.
	DeleteAppInstanceModels []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels `json:"DeleteAppInstanceModels,omitempty" xml:"DeleteAppInstanceModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponseBody) GetDeleteAppInstanceModels() []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	return s.DeleteAppInstanceModels
}

func (s *DeleteAppInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppInstancesResponseBody) SetDeleteAppInstanceModels(v []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels) *DeleteAppInstancesResponseBody {
	s.DeleteAppInstanceModels = v
	return s
}

func (s *DeleteAppInstancesResponseBody) SetRequestId(v string) *DeleteAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteAppInstancesResponseBodyDeleteAppInstanceModels struct {
	// The ID of the application instance.
	//
	// example:
	//
	// ai-gbuea*****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Specifies whether the application instance is deleted.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppInstancesResponseBodyDeleteAppInstanceModels) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GetCode() *string {
	return s.Code
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetAppInstanceId(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetCode(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Code = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetMessage(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Message = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetSuccess(v bool) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Success = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) Validate() error {
	return dara.Validate(s)
}
