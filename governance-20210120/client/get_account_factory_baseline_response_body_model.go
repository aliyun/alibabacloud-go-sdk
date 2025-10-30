// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountFactoryBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v string) *GetAccountFactoryBaselineResponseBody
	GetBaselineId() *string
	SetBaselineItems(v []*GetAccountFactoryBaselineResponseBodyBaselineItems) *GetAccountFactoryBaselineResponseBody
	GetBaselineItems() []*GetAccountFactoryBaselineResponseBodyBaselineItems
	SetBaselineName(v string) *GetAccountFactoryBaselineResponseBody
	GetBaselineName() *string
	SetCreateTime(v string) *GetAccountFactoryBaselineResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetAccountFactoryBaselineResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetAccountFactoryBaselineResponseBody
	GetRequestId() *string
	SetType(v string) *GetAccountFactoryBaselineResponseBody
	GetType() *string
	SetUpdateTime(v string) *GetAccountFactoryBaselineResponseBody
	GetUpdateTime() *string
}

type GetAccountFactoryBaselineResponseBody struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp16ae2k8a3yo3d*****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline items.
	BaselineItems []*GetAccountFactoryBaselineResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The name of the baseline.
	//
	// example:
	//
	// Default
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The time when the baseline was created.
	//
	// example:
	//
	// 2022-11-28T00:46:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the baseline.
	//
	// example:
	//
	// Default baseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60D54503-F1F6-51B6-B6FA-A70CBDA2A68C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the baseline. Valid values:
	//
	// 	- System: default baseline.
	//
	// 	- Custom: custom baseline.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the baseline was updated.
	//
	// example:
	//
	// 2022-11-02T01:00:07Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAccountFactoryBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineResponseBody) GetBaselineId() *string {
	return s.BaselineId
}

func (s *GetAccountFactoryBaselineResponseBody) GetBaselineItems() []*GetAccountFactoryBaselineResponseBodyBaselineItems {
	return s.BaselineItems
}

func (s *GetAccountFactoryBaselineResponseBody) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetAccountFactoryBaselineResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAccountFactoryBaselineResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetAccountFactoryBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountFactoryBaselineResponseBody) GetType() *string {
	return s.Type
}

func (s *GetAccountFactoryBaselineResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAccountFactoryBaselineResponseBody) SetBaselineId(v string) *GetAccountFactoryBaselineResponseBody {
	s.BaselineId = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetBaselineItems(v []*GetAccountFactoryBaselineResponseBodyBaselineItems) *GetAccountFactoryBaselineResponseBody {
	s.BaselineItems = v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetBaselineName(v string) *GetAccountFactoryBaselineResponseBody {
	s.BaselineName = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetCreateTime(v string) *GetAccountFactoryBaselineResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetDescription(v string) *GetAccountFactoryBaselineResponseBody {
	s.Description = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetRequestId(v string) *GetAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetType(v string) *GetAccountFactoryBaselineResponseBody {
	s.Type = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetUpdateTime(v string) *GetAccountFactoryBaselineResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) Validate() error {
	if s.BaselineItems != nil {
		for _, item := range s.BaselineItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAccountFactoryBaselineResponseBodyBaselineItems struct {
	// The configuration of the baseline item.
	//
	// The value is a JSON string.
	//
	// example:
	//
	// {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// 1097526274671790
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAccountFactoryBaselineResponseBodyBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s GetAccountFactoryBaselineResponseBodyBaselineItems) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) GetConfig() *string {
	return s.Config
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) GetName() *string {
	return s.Name
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) SetConfig(v string) *GetAccountFactoryBaselineResponseBodyBaselineItems {
	s.Config = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) SetName(v string) *GetAccountFactoryBaselineResponseBodyBaselineItems {
	s.Name = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) SetVersion(v string) *GetAccountFactoryBaselineResponseBodyBaselineItems {
	s.Version = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) Validate() error {
	return dara.Validate(s)
}
