// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateClusterResponseBody
	GetCode() *int32
	SetData(v *CreateClusterResponseBodyData) *CreateClusterResponseBody
	GetData() *CreateClusterResponseBodyData
	SetErrorCode(v string) *CreateClusterResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateClusterResponseBody
	GetSuccess() *bool
}

type CreateClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *CreateClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B8733786-C045-59F1-8D79-99A52863F62D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateClusterResponseBody) GetData() *CreateClusterResponseBodyData {
	return s.Data
}

func (s *CreateClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateClusterResponseBody) SetCode(v int32) *CreateClusterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClusterResponseBody) SetData(v *CreateClusterResponseBodyData) *CreateClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateClusterResponseBody) SetErrorCode(v string) *CreateClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateClusterResponseBody) SetMessage(v string) *CreateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateClusterResponseBodyData struct {
	// example:
	//
	// xxljob-b21969c2309
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 229317760970086
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateClusterResponseBodyData) SetClusterId(v string) *CreateClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBodyData) SetOrderId(v int64) *CreateClusterResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
