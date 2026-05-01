// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveResourceGroupModelTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*RemoveResourceGroupModelTemplateResponseBodyData) *RemoveResourceGroupModelTemplateResponseBody
	GetData() []*RemoveResourceGroupModelTemplateResponseBodyData
	SetRequestId(v string) *RemoveResourceGroupModelTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *RemoveResourceGroupModelTemplateResponseBody
	GetTotalCount() *int32
}

type RemoveResourceGroupModelTemplateResponseBody struct {
	Data []*RemoveResourceGroupModelTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RemoveResourceGroupModelTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveResourceGroupModelTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveResourceGroupModelTemplateResponseBody) GetData() []*RemoveResourceGroupModelTemplateResponseBodyData {
	return s.Data
}

func (s *RemoveResourceGroupModelTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveResourceGroupModelTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RemoveResourceGroupModelTemplateResponseBody) SetData(v []*RemoveResourceGroupModelTemplateResponseBodyData) *RemoveResourceGroupModelTemplateResponseBody {
	s.Data = v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBody) SetRequestId(v string) *RemoveResourceGroupModelTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBody) SetTotalCount(v int32) *RemoveResourceGroupModelTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBody) Validate() error {
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

type RemoveResourceGroupModelTemplateResponseBodyData struct {
	// example:
	//
	// InvalidParameter.resourceGroupIds
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The parameter resourceGroupIds is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveResourceGroupModelTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveResourceGroupModelTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) SetCode(v string) *RemoveResourceGroupModelTemplateResponseBodyData {
	s.Code = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) SetMessage(v string) *RemoveResourceGroupModelTemplateResponseBodyData {
	s.Message = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) SetResourceGroupId(v string) *RemoveResourceGroupModelTemplateResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) SetSuccess(v bool) *RemoveResourceGroupModelTemplateResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
