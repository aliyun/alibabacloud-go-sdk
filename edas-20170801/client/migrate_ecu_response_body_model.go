// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateEcuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MigrateEcuResponseBody
	GetCode() *int32
	SetData(v string) *MigrateEcuResponseBody
	GetData() *string
	SetMessage(v string) *MigrateEcuResponseBody
	GetMessage() *string
	SetRequestId(v string) *MigrateEcuResponseBody
	GetRequestId() *string
}

type MigrateEcuResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04B0ABAF-95F2-42B6-A7B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateEcuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateEcuResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateEcuResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MigrateEcuResponseBody) GetData() *string {
	return s.Data
}

func (s *MigrateEcuResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateEcuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateEcuResponseBody) SetCode(v int32) *MigrateEcuResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateEcuResponseBody) SetData(v string) *MigrateEcuResponseBody {
	s.Data = &v
	return s
}

func (s *MigrateEcuResponseBody) SetMessage(v string) *MigrateEcuResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateEcuResponseBody) SetRequestId(v string) *MigrateEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateEcuResponseBody) Validate() error {
	return dara.Validate(s)
}
