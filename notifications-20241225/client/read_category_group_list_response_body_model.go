// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadCategoryGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadCategoryGroupListResponseBody
	GetCode() *string
	SetData(v []*ReadCategoryGroupListResponseBodyData) *ReadCategoryGroupListResponseBody
	GetData() []*ReadCategoryGroupListResponseBodyData
	SetMessage(v string) *ReadCategoryGroupListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadCategoryGroupListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadCategoryGroupListResponseBody
	GetSuccess() *bool
}

type ReadCategoryGroupListResponseBody struct {
	// The error code returned by the system. For more information about error codes, see Error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	Data []*ReadCategoryGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned when the call fails.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// /
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadCategoryGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadCategoryGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadCategoryGroupListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadCategoryGroupListResponseBody) GetData() []*ReadCategoryGroupListResponseBodyData {
	return s.Data
}

func (s *ReadCategoryGroupListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadCategoryGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadCategoryGroupListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadCategoryGroupListResponseBody) SetCode(v string) *ReadCategoryGroupListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadCategoryGroupListResponseBody) SetData(v []*ReadCategoryGroupListResponseBodyData) *ReadCategoryGroupListResponseBody {
	s.Data = v
	return s
}

func (s *ReadCategoryGroupListResponseBody) SetMessage(v string) *ReadCategoryGroupListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadCategoryGroupListResponseBody) SetRequestId(v string) *ReadCategoryGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadCategoryGroupListResponseBody) SetSuccess(v bool) *ReadCategoryGroupListResponseBody {
	s.Success = &v
	return s
}

func (s *ReadCategoryGroupListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadCategoryGroupListResponseBodyData struct {
	// The category group code.
	//
	// example:
	//
	// prod_msg
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The group name.
	//
	// example:
	//
	// 产品消息
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ReadCategoryGroupListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadCategoryGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadCategoryGroupListResponseBodyData) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ReadCategoryGroupListResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ReadCategoryGroupListResponseBodyData) SetGroupCode(v string) *ReadCategoryGroupListResponseBodyData {
	s.GroupCode = &v
	return s
}

func (s *ReadCategoryGroupListResponseBodyData) SetGroupName(v string) *ReadCategoryGroupListResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ReadCategoryGroupListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
