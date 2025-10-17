// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyParameterLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeModifyParameterLogResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeModifyParameterLogResponseBody
	GetEngineVersion() *string
	SetItems(v []*DescribeModifyParameterLogResponseBodyItems) *DescribeModifyParameterLogResponseBody
	GetItems() []*DescribeModifyParameterLogResponseBodyItems
	SetRequestId(v string) *DescribeModifyParameterLogResponseBody
	GetRequestId() *string
}

type DescribeModifyParameterLogResponseBody struct {
	// example:
	//
	// polardb_mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 8.0
	EngineVersion *string                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Items         []*DescribeModifyParameterLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModifyParameterLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeModifyParameterLogResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeModifyParameterLogResponseBody) GetItems() []*DescribeModifyParameterLogResponseBodyItems {
	return s.Items
}

func (s *DescribeModifyParameterLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModifyParameterLogResponseBody) SetEngine(v string) *DescribeModifyParameterLogResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetEngineVersion(v string) *DescribeModifyParameterLogResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetItems(v []*DescribeModifyParameterLogResponseBodyItems) *DescribeModifyParameterLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetRequestId(v string) *DescribeModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModifyParameterLogResponseBodyItems struct {
	// example:
	//
	// 2024-10-29T09:31:37Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test01
	NewParameterValue *string `json:"NewParameterValue,omitempty" xml:"NewParameterValue,omitempty"`
	// example:
	//
	// test
	OldParameterValue *string `json:"OldParameterValue,omitempty" xml:"OldParameterValue,omitempty"`
	// example:
	//
	// hz
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// True
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModifyParameterLogResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeModifyParameterLogResponseBodyItems) GetNewParameterValue() *string {
	return s.NewParameterValue
}

func (s *DescribeModifyParameterLogResponseBodyItems) GetOldParameterValue() *string {
	return s.OldParameterValue
}

func (s *DescribeModifyParameterLogResponseBodyItems) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeModifyParameterLogResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeModifyParameterLogResponseBodyItems) SetModifyTime(v string) *DescribeModifyParameterLogResponseBodyItems {
	s.ModifyTime = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItems) SetNewParameterValue(v string) *DescribeModifyParameterLogResponseBodyItems {
	s.NewParameterValue = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItems) SetOldParameterValue(v string) *DescribeModifyParameterLogResponseBodyItems {
	s.OldParameterValue = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItems) SetParameterName(v string) *DescribeModifyParameterLogResponseBodyItems {
	s.ParameterName = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItems) SetStatus(v string) *DescribeModifyParameterLogResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
