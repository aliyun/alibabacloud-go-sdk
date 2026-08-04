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
	// The error code returned. A value of 200 indicates that the call succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return result of invoking this API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of user information returned.
	//
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserByDeviceIdResponseBodyResult struct {
	// The openID corresponding to the user information.
	//
	// example:
	//
	// 0963*0158
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
	// The list of organization IDs and UnionIDs for the user.
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
	if s.UserUnionIds != nil {
		for _, item := range s.UserUnionIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserByDeviceIdResponseBodyResultUserUnionIds struct {
	// The organization ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The user\\"s UnionID.
	//
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
