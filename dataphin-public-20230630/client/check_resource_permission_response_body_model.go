// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckResourcePermissionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CheckResourcePermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckResourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckResourcePermissionResponseBody
	GetRequestId() *string
	SetResourcePermissionList(v []*CheckResourcePermissionResponseBodyResourcePermissionList) *CheckResourcePermissionResponseBody
	GetResourcePermissionList() []*CheckResourcePermissionResponseBodyResourcePermissionList
	SetSuccess(v bool) *CheckResourcePermissionResponseBody
	GetSuccess() *bool
}

type CheckResourcePermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePermissionList []*CheckResourcePermissionResponseBodyResourcePermissionList `json:"ResourcePermissionList,omitempty" xml:"ResourcePermissionList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckResourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckResourcePermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckResourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckResourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckResourcePermissionResponseBody) GetResourcePermissionList() []*CheckResourcePermissionResponseBodyResourcePermissionList {
	return s.ResourcePermissionList
}

func (s *CheckResourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckResourcePermissionResponseBody) SetCode(v string) *CheckResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetHttpStatusCode(v int32) *CheckResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetMessage(v string) *CheckResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetRequestId(v string) *CheckResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetResourcePermissionList(v []*CheckResourcePermissionResponseBodyResourcePermissionList) *CheckResourcePermissionResponseBody {
	s.ResourcePermissionList = v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetSuccess(v bool) *CheckResourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) Validate() error {
	if s.ResourcePermissionList != nil {
		for _, item := range s.ResourcePermissionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckResourcePermissionResponseBodyResourcePermissionList struct {
	// example:
	//
	// true
	HasPermission *bool `json:"HasPermission,omitempty" xml:"HasPermission,omitempty"`
	// example:
	//
	// hadoop.300000806.data_distill.behavior_gameinfor_01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CheckResourcePermissionResponseBodyResourcePermissionList) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionResponseBodyResourcePermissionList) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) GetHasPermission() *bool {
	return s.HasPermission
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) GetResourceId() *string {
	return s.ResourceId
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) SetHasPermission(v bool) *CheckResourcePermissionResponseBodyResourcePermissionList {
	s.HasPermission = &v
	return s
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) SetResourceId(v string) *CheckResourcePermissionResponseBodyResourcePermissionList {
	s.ResourceId = &v
	return s
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) Validate() error {
	return dara.Validate(s)
}
