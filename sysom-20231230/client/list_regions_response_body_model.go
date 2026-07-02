// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
	SetCode(v string) *ListRegionsResponseBody
	GetCode() *string
	SetData(v []*string) *ListRegionsResponseBody
	GetData() []*string
	SetMessage(v string) *ListRegionsResponseBody
	GetMessage() *string
}

type ListRegionsResponseBody struct {
	// The request ID, which can be used for end-to-end diagnostics.
	//
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of regions.
	//
	// example:
	//
	// ["cn-hangzhou", "cn-shengzhen"]
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code description. This value is empty if no error occurs.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRegionsResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetCode(v string) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetData(v []*string) *ListRegionsResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
