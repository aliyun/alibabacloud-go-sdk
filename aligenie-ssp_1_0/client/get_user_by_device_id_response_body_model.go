// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByDeviceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetUserByDeviceIdResponseBody
	GetCode() *int32
	SetMessage(v string) *GetUserByDeviceIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserByDeviceIdResponseBody
	GetRequestId() *string
	SetResult(v *GetUserByDeviceIdResponseBodyResult) *GetUserByDeviceIdResponseBody
	GetResult() *GetUserByDeviceIdResponseBodyResult
}

type GetUserByDeviceIdResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *GetUserByDeviceIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetUserByDeviceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetUserByDeviceIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserByDeviceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserByDeviceIdResponseBody) GetResult() *GetUserByDeviceIdResponseBodyResult {
	return s.Result
}

func (s *GetUserByDeviceIdResponseBody) SetCode(v int32) *GetUserByDeviceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserByDeviceIdResponseBody) SetMessage(v string) *GetUserByDeviceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserByDeviceIdResponseBody) SetRequestId(v string) *GetUserByDeviceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBody) SetResult(v *GetUserByDeviceIdResponseBodyResult) *GetUserByDeviceIdResponseBody {
	s.Result = v
	return s
}

func (s *GetUserByDeviceIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserByDeviceIdResponseBodyResult struct {
	// example:
	//
	// 0963*0158
	UserOpenId   *string                                            `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
	UserUnionIds []*GetUserByDeviceIdResponseBodyResultUserUnionIds `json:"UserUnionIds,omitempty" xml:"UserUnionIds,omitempty" type:"Repeated"`
}

func (s GetUserByDeviceIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponseBodyResult) GetUserOpenId() *string {
	return s.UserOpenId
}

func (s *GetUserByDeviceIdResponseBodyResult) GetUserUnionIds() []*GetUserByDeviceIdResponseBodyResultUserUnionIds {
	return s.UserUnionIds
}

func (s *GetUserByDeviceIdResponseBodyResult) SetUserOpenId(v string) *GetUserByDeviceIdResponseBodyResult {
	s.UserOpenId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBodyResult) SetUserUnionIds(v []*GetUserByDeviceIdResponseBodyResultUserUnionIds) *GetUserByDeviceIdResponseBodyResult {
	s.UserUnionIds = v
	return s
}

func (s *GetUserByDeviceIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetUserByDeviceIdResponseBodyResultUserUnionIds struct {
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 1553*B0C3
	UserUnionId *string `json:"UserUnionId,omitempty" xml:"UserUnionId,omitempty"`
}

func (s GetUserByDeviceIdResponseBodyResultUserUnionIds) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdResponseBodyResultUserUnionIds) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) GetUserUnionId() *string {
	return s.UserUnionId
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) SetOrganizationId(v string) *GetUserByDeviceIdResponseBodyResultUserUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) SetUserUnionId(v string) *GetUserByDeviceIdResponseBodyResultUserUnionIds {
	s.UserUnionId = &v
	return s
}

func (s *GetUserByDeviceIdResponseBodyResultUserUnionIds) Validate() error {
	return dara.Validate(s)
}
