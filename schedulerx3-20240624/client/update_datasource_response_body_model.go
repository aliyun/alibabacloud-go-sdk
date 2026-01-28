// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateDatasourceResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateDatasourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDatasourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDatasourceResponseBody
	GetSuccess() *bool
}

type UpdateDatasourceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter format error
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

func (s UpdateDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateDatasourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDatasourceResponseBody) SetCode(v int32) *UpdateDatasourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDatasourceResponseBody) SetMessage(v string) *UpdateDatasourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDatasourceResponseBody) SetRequestId(v string) *UpdateDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasourceResponseBody) SetSuccess(v bool) *UpdateDatasourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
