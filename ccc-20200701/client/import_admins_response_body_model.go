// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAdminsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportAdminsResponseBody
	GetCode() *string
	SetData(v []*ImportAdminsResponseBodyData) *ImportAdminsResponseBody
	GetData() []*ImportAdminsResponseBodyData
	SetHttpStatusCode(v int32) *ImportAdminsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportAdminsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportAdminsResponseBody
	GetRequestId() *string
}

type ImportAdminsResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ImportAdminsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CDB5C94-ACFB-59B5-85FE-C1DAF8049C63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportAdminsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportAdminsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAdminsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportAdminsResponseBody) GetData() []*ImportAdminsResponseBodyData {
	return s.Data
}

func (s *ImportAdminsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportAdminsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportAdminsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportAdminsResponseBody) SetCode(v string) *ImportAdminsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportAdminsResponseBody) SetData(v []*ImportAdminsResponseBodyData) *ImportAdminsResponseBody {
	s.Data = v
	return s
}

func (s *ImportAdminsResponseBody) SetHttpStatusCode(v int32) *ImportAdminsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportAdminsResponseBody) SetMessage(v string) *ImportAdminsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportAdminsResponseBody) SetRequestId(v string) *ImportAdminsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportAdminsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImportAdminsResponseBodyData struct {
	// example:
	//
	// 8021****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 26972543893791****
	RamId *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ImportAdminsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportAdminsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportAdminsResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ImportAdminsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportAdminsResponseBodyData) GetRamId() *string {
	return s.RamId
}

func (s *ImportAdminsResponseBodyData) GetRoleId() *string {
	return s.RoleId
}

func (s *ImportAdminsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ImportAdminsResponseBodyData) SetExtension(v string) *ImportAdminsResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetInstanceId(v string) *ImportAdminsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetRamId(v string) *ImportAdminsResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetRoleId(v string) *ImportAdminsResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetUserId(v string) *ImportAdminsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
