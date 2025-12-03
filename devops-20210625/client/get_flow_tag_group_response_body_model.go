// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetFlowTagGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFlowTagGroupResponseBody
	GetErrorMessage() *string
	SetFlowTagGroup(v *GetFlowTagGroupResponseBodyFlowTagGroup) *GetFlowTagGroupResponseBody
	GetFlowTagGroup() *GetFlowTagGroupResponseBodyFlowTagGroup
	SetRequestId(v string) *GetFlowTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFlowTagGroupResponseBody
	GetSuccess() *bool
}

type GetFlowTagGroupResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FlowTagGroup *GetFlowTagGroupResponseBodyFlowTagGroup `json:"flowTagGroup,omitempty" xml:"flowTagGroup,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFlowTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFlowTagGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFlowTagGroupResponseBody) GetFlowTagGroup() *GetFlowTagGroupResponseBodyFlowTagGroup {
	return s.FlowTagGroup
}

func (s *GetFlowTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlowTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFlowTagGroupResponseBody) SetErrorCode(v string) *GetFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetErrorMessage(v string) *GetFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetFlowTagGroup(v *GetFlowTagGroupResponseBodyFlowTagGroup) *GetFlowTagGroupResponseBody {
	s.FlowTagGroup = v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetRequestId(v string) *GetFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetSuccess(v bool) *GetFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) Validate() error {
	if s.FlowTagGroup != nil {
		if err := s.FlowTagGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFlowTagGroupResponseBodyFlowTagGroup struct {
	// example:
	//
	// 1111111111111
	CreatorAccountId *string                                               `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	FlowTagList      []*GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList `json:"flowTagList,omitempty" xml:"flowTagList,omitempty" type:"Repeated"`
	// example:
	//
	// 1111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1111111111111
	ModiferAccountId *string `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetFlowTagGroupResponseBodyFlowTagGroup) String() string {
	return dara.Prettify(s)
}

func (s GetFlowTagGroupResponseBodyFlowTagGroup) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) GetFlowTagList() []*GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	return s.FlowTagList
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) GetId() *int64 {
	return s.Id
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) GetModiferAccountId() *string {
	return s.ModiferAccountId
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) GetName() *string {
	return s.Name
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetCreatorAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.CreatorAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetFlowTagList(v []*GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.FlowTagList = v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetId(v int64) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.Id = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetModiferAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.ModiferAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetName(v string) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.Name = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) Validate() error {
	if s.FlowTagList != nil {
		for _, item := range s.FlowTagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList struct {
	// example:
	//
	// #1F9AEF
	Color *string `json:"color,omitempty" xml:"color,omitempty"`
	// example:
	//
	// 1111111111111
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1111111111111
	ModiferAccountId *string `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) String() string {
	return dara.Prettify(s)
}

func (s GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GetColor() *string {
	return s.Color
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GetId() *int64 {
	return s.Id
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GetModiferAccountId() *string {
	return s.ModiferAccountId
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GetName() *string {
	return s.Name
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetColor(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.Color = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetCreatorAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.CreatorAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetId(v int64) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.Id = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetModiferAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.ModiferAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetName(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.Name = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) Validate() error {
	return dara.Validate(s)
}
