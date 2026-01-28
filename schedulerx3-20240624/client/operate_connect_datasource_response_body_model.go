// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateConnectDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateConnectDatasourceResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateConnectDatasourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateConnectDatasourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateConnectDatasourceResponseBody
	GetSuccess() *bool
}

type OperateConnectDatasourceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BAC1ADB5-EEB5-5834-93D8-522E067AF8D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateConnectDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateConnectDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *OperateConnectDatasourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateConnectDatasourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateConnectDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateConnectDatasourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateConnectDatasourceResponseBody) SetCode(v int32) *OperateConnectDatasourceResponseBody {
	s.Code = &v
	return s
}

func (s *OperateConnectDatasourceResponseBody) SetMessage(v string) *OperateConnectDatasourceResponseBody {
	s.Message = &v
	return s
}

func (s *OperateConnectDatasourceResponseBody) SetRequestId(v string) *OperateConnectDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateConnectDatasourceResponseBody) SetSuccess(v bool) *OperateConnectDatasourceResponseBody {
	s.Success = &v
	return s
}

func (s *OperateConnectDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
