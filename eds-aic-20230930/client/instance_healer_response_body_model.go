// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceHealerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceHealerModel(v *InstanceHealerResponseBodyInstanceHealerModel) *InstanceHealerResponseBody
	GetInstanceHealerModel() *InstanceHealerResponseBodyInstanceHealerModel
	SetRequestId(v string) *InstanceHealerResponseBody
	GetRequestId() *string
}

type InstanceHealerResponseBody struct {
	InstanceHealerModel *InstanceHealerResponseBodyInstanceHealerModel `json:"InstanceHealerModel,omitempty" xml:"InstanceHealerModel,omitempty" type:"Struct"`
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstanceHealerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstanceHealerResponseBody) GoString() string {
	return s.String()
}

func (s *InstanceHealerResponseBody) GetInstanceHealerModel() *InstanceHealerResponseBodyInstanceHealerModel {
	return s.InstanceHealerModel
}

func (s *InstanceHealerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstanceHealerResponseBody) SetInstanceHealerModel(v *InstanceHealerResponseBodyInstanceHealerModel) *InstanceHealerResponseBody {
	s.InstanceHealerModel = v
	return s
}

func (s *InstanceHealerResponseBody) SetRequestId(v string) *InstanceHealerResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstanceHealerResponseBody) Validate() error {
	if s.InstanceHealerModel != nil {
		if err := s.InstanceHealerModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstanceHealerResponseBodyInstanceHealerModel struct {
	// example:
	//
	// True
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InstanceHealerResponseBodyInstanceHealerModel) String() string {
	return dara.Prettify(s)
}

func (s InstanceHealerResponseBodyInstanceHealerModel) GoString() string {
	return s.String()
}

func (s *InstanceHealerResponseBodyInstanceHealerModel) GetResult() *string {
	return s.Result
}

func (s *InstanceHealerResponseBodyInstanceHealerModel) SetResult(v string) *InstanceHealerResponseBodyInstanceHealerModel {
	s.Result = &v
	return s
}

func (s *InstanceHealerResponseBodyInstanceHealerModel) Validate() error {
	return dara.Validate(s)
}
