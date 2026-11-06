// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupModel(v *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) *ModifyAppInstanceGroupChargeTypeResponseBody
	GetAppInstanceGroupModel() *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel
	SetRequestId(v string) *ModifyAppInstanceGroupChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyAppInstanceGroupChargeTypeResponseBody struct {
	// The task information for this billing method conversion.
	AppInstanceGroupModel *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel `json:"AppInstanceGroupModel,omitempty" xml:"AppInstanceGroupModel,omitempty" type:"Struct"`
	// The unique ID of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppInstanceGroupChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBody) GetAppInstanceGroupModel() *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel {
	return s.AppInstanceGroupModel
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBody) SetAppInstanceGroupModel(v *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) *ModifyAppInstanceGroupChargeTypeResponseBody {
	s.AppInstanceGroupModel = v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBody) SetRequestId(v string) *ModifyAppInstanceGroupChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBody) Validate() error {
	if s.AppInstanceGroupModel != nil {
		if err := s.AppInstanceGroupModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel struct {
	// The ID of the billing conversion order task, used to associate with this conversion. It is returned only when a related order task is required and created. It does not indicate that the payment is completed or the new billing method has taken effect.
	//
	// example:
	//
	// task-123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) SetTaskId(v string) *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel {
	s.TaskId = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeResponseBodyAppInstanceGroupModel) Validate() error {
	return dara.Validate(s)
}
