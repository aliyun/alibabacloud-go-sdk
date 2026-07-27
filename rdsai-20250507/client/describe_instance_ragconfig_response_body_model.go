// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRAGConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceRAGConfigResponseBody
	GetBranchName() *string
	SetConfigList(v []*DescribeInstanceRAGConfigResponseBodyConfigList) *DescribeInstanceRAGConfigResponseBody
	GetConfigList() []*DescribeInstanceRAGConfigResponseBodyConfigList
	SetInstanceName(v string) *DescribeInstanceRAGConfigResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DescribeInstanceRAGConfigResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DescribeInstanceRAGConfigResponseBody
	GetStatus() *bool
}

type DescribeInstanceRAGConfigResponseBody struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The list of RAG Agent configurations.
	ConfigList []*DescribeInstanceRAGConfigResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The RAG Agent status. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceRAGConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRAGConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRAGConfigResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceRAGConfigResponseBody) GetConfigList() []*DescribeInstanceRAGConfigResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeInstanceRAGConfigResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceRAGConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRAGConfigResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DescribeInstanceRAGConfigResponseBody) SetBranchName(v string) *DescribeInstanceRAGConfigResponseBody {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBody) SetConfigList(v []*DescribeInstanceRAGConfigResponseBodyConfigList) *DescribeInstanceRAGConfigResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBody) SetInstanceName(v string) *DescribeInstanceRAGConfigResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBody) SetRequestId(v string) *DescribeInstanceRAGConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBody) SetStatus(v bool) *DescribeInstanceRAGConfigResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBody) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRAGConfigResponseBodyConfigList struct {
	// The name of the configuration item.
	//
	// example:
	//
	// LLM_MODEL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// qwen-flash
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceRAGConfigResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRAGConfigResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRAGConfigResponseBodyConfigList) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceRAGConfigResponseBodyConfigList) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceRAGConfigResponseBodyConfigList) SetName(v string) *DescribeInstanceRAGConfigResponseBodyConfigList {
	s.Name = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBodyConfigList) SetValue(v string) *DescribeInstanceRAGConfigResponseBodyConfigList {
	s.Value = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
