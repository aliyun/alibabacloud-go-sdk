// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiccsRobotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAiccsRobotResponseBody
	GetCode() *string
	SetData(v []*ListAiccsRobotResponseBodyData) *ListAiccsRobotResponseBody
	GetData() []*ListAiccsRobotResponseBodyData
	SetMessage(v string) *ListAiccsRobotResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAiccsRobotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAiccsRobotResponseBody
	GetSuccess() *bool
}

type ListAiccsRobotResponseBody struct {
	// Request status code. A return value of OK indicates that the request Succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Robot scripts.
	Data []*ListAiccsRobotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation Succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAiccsRobotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiccsRobotResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAiccsRobotResponseBody) GetData() []*ListAiccsRobotResponseBodyData {
	return s.Data
}

func (s *ListAiccsRobotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAiccsRobotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiccsRobotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAiccsRobotResponseBody) SetCode(v string) *ListAiccsRobotResponseBody {
	s.Code = &v
	return s
}

func (s *ListAiccsRobotResponseBody) SetData(v []*ListAiccsRobotResponseBodyData) *ListAiccsRobotResponseBody {
	s.Data = v
	return s
}

func (s *ListAiccsRobotResponseBody) SetMessage(v string) *ListAiccsRobotResponseBody {
	s.Message = &v
	return s
}

func (s *ListAiccsRobotResponseBody) SetRequestId(v string) *ListAiccsRobotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiccsRobotResponseBody) SetSuccess(v bool) *ListAiccsRobotResponseBody {
	s.Success = &v
	return s
}

func (s *ListAiccsRobotResponseBody) Validate() error {
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

type ListAiccsRobotResponseBodyData struct {
	// Associated industry.
	//
	// example:
	//
	// 房地产
	AtProfession *string `json:"AtProfession,omitempty" xml:"AtProfession,omitempty"`
	// Associated business.
	//
	// example:
	//
	// 新房销售
	AtSence *string `json:"AtSence,omitempty" xml:"AtSence,omitempty"`
	// Robot ID.
	//
	// example:
	//
	// 12****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Robot name.
	//
	// example:
	//
	// 测试机器人
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// Robot type.
	//
	// example:
	//
	// CUSTOM
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s ListAiccsRobotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAiccsRobotResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotResponseBodyData) GetAtProfession() *string {
	return s.AtProfession
}

func (s *ListAiccsRobotResponseBodyData) GetAtSence() *string {
	return s.AtSence
}

func (s *ListAiccsRobotResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAiccsRobotResponseBodyData) GetRobotName() *string {
	return s.RobotName
}

func (s *ListAiccsRobotResponseBodyData) GetRobotType() *string {
	return s.RobotType
}

func (s *ListAiccsRobotResponseBodyData) SetAtProfession(v string) *ListAiccsRobotResponseBodyData {
	s.AtProfession = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetAtSence(v string) *ListAiccsRobotResponseBodyData {
	s.AtSence = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetId(v int64) *ListAiccsRobotResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetRobotName(v string) *ListAiccsRobotResponseBodyData {
	s.RobotName = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) SetRobotType(v string) *ListAiccsRobotResponseBodyData {
	s.RobotType = &v
	return s
}

func (s *ListAiccsRobotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
