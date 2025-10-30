// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareCopyRuleVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompareCopyRuleVariableResponseBody
	GetRequestId() *string
	SetResultObject(v *CompareCopyRuleVariableResponseBodyResultObject) *CompareCopyRuleVariableResponseBody
	GetResultObject() *CompareCopyRuleVariableResponseBodyResultObject
}

type CompareCopyRuleVariableResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	ResultObject *CompareCopyRuleVariableResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CompareCopyRuleVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareCopyRuleVariableResponseBody) GetResultObject() *CompareCopyRuleVariableResponseBodyResultObject {
	return s.ResultObject
}

func (s *CompareCopyRuleVariableResponseBody) SetRequestId(v string) *CompareCopyRuleVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBody) SetResultObject(v *CompareCopyRuleVariableResponseBodyResultObject) *CompareCopyRuleVariableResponseBody {
	s.ResultObject = v
	return s
}

func (s *CompareCopyRuleVariableResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareCopyRuleVariableResponseBodyResultObject struct {
	// Cumulative variable list
	CustVariableList []*CompareCopyRuleVariableResponseBodyResultObjectCustVariableList `json:"CustVariableList,omitempty" xml:"CustVariableList,omitempty" type:"Repeated"`
	// Event field variables
	EventVariableList []*CompareCopyRuleVariableResponseBodyResultObjectEventVariableList `json:"EventVariableList,omitempty" xml:"EventVariableList,omitempty" type:"Repeated"`
	// Custom variable list
	ExpressionVariableList []*CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList `json:"ExpressionVariableList,omitempty" xml:"ExpressionVariableList,omitempty" type:"Repeated"`
	// Name list variables
	NameListVariableList []*CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList `json:"NameListVariableList,omitempty" xml:"NameListVariableList,omitempty" type:"Repeated"`
	// Custom Query Variable List
	QueryExpressionVariableList []*CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList `json:"QueryExpressionVariableList,omitempty" xml:"QueryExpressionVariableList,omitempty" type:"Repeated"`
	// System variable list
	SystemVariableList []*CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList `json:"SystemVariableList,omitempty" xml:"SystemVariableList,omitempty" type:"Repeated"`
}

func (s CompareCopyRuleVariableResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) GetCustVariableList() []*CompareCopyRuleVariableResponseBodyResultObjectCustVariableList {
	return s.CustVariableList
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) GetEventVariableList() []*CompareCopyRuleVariableResponseBodyResultObjectEventVariableList {
	return s.EventVariableList
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) GetExpressionVariableList() []*CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList {
	return s.ExpressionVariableList
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) GetNameListVariableList() []*CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList {
	return s.NameListVariableList
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) GetQueryExpressionVariableList() []*CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList {
	return s.QueryExpressionVariableList
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) GetSystemVariableList() []*CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList {
	return s.SystemVariableList
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) SetCustVariableList(v []*CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) *CompareCopyRuleVariableResponseBodyResultObject {
	s.CustVariableList = v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) SetEventVariableList(v []*CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) *CompareCopyRuleVariableResponseBodyResultObject {
	s.EventVariableList = v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) SetExpressionVariableList(v []*CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) *CompareCopyRuleVariableResponseBodyResultObject {
	s.ExpressionVariableList = v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) SetNameListVariableList(v []*CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) *CompareCopyRuleVariableResponseBodyResultObject {
	s.NameListVariableList = v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) SetQueryExpressionVariableList(v []*CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) *CompareCopyRuleVariableResponseBodyResultObject {
	s.QueryExpressionVariableList = v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) SetSystemVariableList(v []*CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) *CompareCopyRuleVariableResponseBodyResultObject {
	s.SystemVariableList = v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObject) Validate() error {
	if s.CustVariableList != nil {
		for _, item := range s.CustVariableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EventVariableList != nil {
		for _, item := range s.EventVariableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExpressionVariableList != nil {
		for _, item := range s.ExpressionVariableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NameListVariableList != nil {
		for _, item := range s.NameListVariableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryExpressionVariableList != nil {
		for _, item := range s.QueryExpressionVariableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemVariableList != nil {
		for _, item := range s.SystemVariableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CompareCopyRuleVariableResponseBodyResultObjectCustVariableList struct {
	// Description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 1483
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Variable name
	//
	// example:
	//
	// dxkkLw8fe18
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Variable type
	//
	// example:
	//
	// SELF_VELOCITY
	OutType *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	// Title
	//
	// example:
	//
	// 近1天账户登录设备数_bk4
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) GetDescription() *string {
	return s.Description
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) GetId() *int64 {
	return s.Id
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) GetName() *string {
	return s.Name
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) GetOutType() *string {
	return s.OutType
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) GetTitle() *string {
	return s.Title
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) SetDescription(v string) *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList {
	s.Description = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) SetId(v int64) *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList {
	s.Id = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) SetName(v string) *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList {
	s.Name = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) SetOutType(v string) *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList {
	s.OutType = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) SetTitle(v string) *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList {
	s.Title = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectCustVariableList) Validate() error {
	return dara.Validate(s)
}

type CompareCopyRuleVariableResponseBodyResultObjectEventVariableList struct {
	// Description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Variable id
	//
	// example:
	//
	// 375
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Variable name
	//
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Variable type
	//
	// example:
	//
	// NATIVE
	OutType *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	// Title
	//
	// example:
	//
	// 年龄
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) GetDescription() *string {
	return s.Description
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) GetId() *int64 {
	return s.Id
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) GetName() *string {
	return s.Name
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) GetOutType() *string {
	return s.OutType
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) GetTitle() *string {
	return s.Title
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) SetDescription(v string) *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList {
	s.Description = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) SetId(v int64) *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList {
	s.Id = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) SetName(v string) *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList {
	s.Name = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) SetOutType(v string) *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList {
	s.OutType = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) SetTitle(v string) *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList {
	s.Title = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectEventVariableList) Validate() error {
	return dara.Validate(s)
}

type CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList struct {
	// Description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 2540
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Variable name
	//
	// example:
	//
	// ex_XheC9A382fe7
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Variable Type
	//
	// example:
	//
	// EXPRESSION
	OutType *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	// Title
	//
	// example:
	//
	// 手机号前7位
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) GetDescription() *string {
	return s.Description
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) GetId() *int64 {
	return s.Id
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) GetName() *string {
	return s.Name
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) GetOutType() *string {
	return s.OutType
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) GetTitle() *string {
	return s.Title
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) SetDescription(v string) *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList {
	s.Description = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) SetId(v int64) *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList {
	s.Id = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) SetName(v string) *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList {
	s.Name = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) SetOutType(v string) *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList {
	s.OutType = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) SetTitle(v string) *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList {
	s.Title = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectExpressionVariableList) Validate() error {
	return dara.Validate(s)
}

type CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList struct {
	// Description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Variable id
	//
	// example:
	//
	// 1987
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Variable name
	//
	// example:
	//
	// nl_UN8otElLb490
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Variable type
	//
	// example:
	//
	// NAME_LIST
	OutType *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	// Title
	//
	// example:
	//
	// 白名单
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) GetDescription() *string {
	return s.Description
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) GetId() *int64 {
	return s.Id
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) GetName() *string {
	return s.Name
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) GetOutType() *string {
	return s.OutType
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) GetTitle() *string {
	return s.Title
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) SetDescription(v string) *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList {
	s.Description = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) SetId(v int64) *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList {
	s.Id = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) SetName(v string) *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList {
	s.Name = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) SetOutType(v string) *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList {
	s.OutType = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) SetTitle(v string) *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList {
	s.Title = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectNameListVariableList) Validate() error {
	return dara.Validate(s)
}

type CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList struct {
	// Description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 217
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Variable Name
	//
	// example:
	//
	// ex_vcaoii1697
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Variable Type
	//
	// example:
	//
	// QUERY_EXPRESSION
	OutType *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	// Title
	//
	// example:
	//
	// 获取年龄
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) GetDescription() *string {
	return s.Description
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) GetId() *int64 {
	return s.Id
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) GetName() *string {
	return s.Name
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) GetOutType() *string {
	return s.OutType
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) GetTitle() *string {
	return s.Title
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) SetDescription(v string) *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList {
	s.Description = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) SetId(v int64) *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList {
	s.Id = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) SetName(v string) *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList {
	s.Name = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) SetOutType(v string) *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList {
	s.OutType = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) SetTitle(v string) *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList {
	s.Title = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectQueryExpressionVariableList) Validate() error {
	return dara.Validate(s)
}

type CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList struct {
	// Description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 1003
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Variable name
	//
	// example:
	//
	// sl_S02sHLFT7818
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Variable type
	//
	// example:
	//
	// SELF_BIND
	OutType *string `json:"OutType,omitempty" xml:"OutType,omitempty"`
	// Title
	//
	// example:
	//
	// 根据ip地址得到评分
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) GetDescription() *string {
	return s.Description
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) GetId() *int64 {
	return s.Id
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) GetName() *string {
	return s.Name
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) GetOutType() *string {
	return s.OutType
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) GetTitle() *string {
	return s.Title
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) SetDescription(v string) *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList {
	s.Description = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) SetId(v int64) *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList {
	s.Id = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) SetName(v string) *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList {
	s.Name = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) SetOutType(v string) *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList {
	s.OutType = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) SetTitle(v string) *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList {
	s.Title = &v
	return s
}

func (s *CompareCopyRuleVariableResponseBodyResultObjectSystemVariableList) Validate() error {
	return dara.Validate(s)
}
