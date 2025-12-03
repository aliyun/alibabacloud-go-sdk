// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowTagGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListFlowTagGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListFlowTagGroupsResponseBody
	GetErrorMessage() *string
	SetFlowTagGroups(v []*ListFlowTagGroupsResponseBodyFlowTagGroups) *ListFlowTagGroupsResponseBody
	GetFlowTagGroups() []*ListFlowTagGroupsResponseBodyFlowTagGroups
	SetRequestId(v string) *ListFlowTagGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlowTagGroupsResponseBody
	GetSuccess() *bool
}

type ListFlowTagGroupsResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage  *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FlowTagGroups []*ListFlowTagGroupsResponseBodyFlowTagGroups `json:"flowTagGroups,omitempty" xml:"flowTagGroups,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListFlowTagGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowTagGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowTagGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListFlowTagGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListFlowTagGroupsResponseBody) GetFlowTagGroups() []*ListFlowTagGroupsResponseBodyFlowTagGroups {
	return s.FlowTagGroups
}

func (s *ListFlowTagGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowTagGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlowTagGroupsResponseBody) SetErrorCode(v string) *ListFlowTagGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetErrorMessage(v string) *ListFlowTagGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetFlowTagGroups(v []*ListFlowTagGroupsResponseBodyFlowTagGroups) *ListFlowTagGroupsResponseBody {
	s.FlowTagGroups = v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetRequestId(v string) *ListFlowTagGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetSuccess(v bool) *ListFlowTagGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) Validate() error {
	if s.FlowTagGroups != nil {
		for _, item := range s.FlowTagGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowTagGroupsResponseBodyFlowTagGroups struct {
	// example:
	//
	// 111111111
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 11111111
	ModiferAccountId *string `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	// example:
	//
	// 标签名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListFlowTagGroupsResponseBodyFlowTagGroups) String() string {
	return dara.Prettify(s)
}

func (s ListFlowTagGroupsResponseBodyFlowTagGroups) GoString() string {
	return s.String()
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) GetId() *int64 {
	return s.Id
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) GetModiferAccountId() *string {
	return s.ModiferAccountId
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) GetName() *string {
	return s.Name
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetCreatorAccountId(v string) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.CreatorAccountId = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetId(v int64) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.Id = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetModiferAccountId(v string) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.ModiferAccountId = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetName(v string) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.Name = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) Validate() error {
	return dara.Validate(s)
}
