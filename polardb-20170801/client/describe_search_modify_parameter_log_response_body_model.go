// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchModifyParameterLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeSearchModifyParameterLogResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeSearchModifyParameterLogResponseBody
	GetEngineVersion() *string
	SetItems(v []*DescribeSearchModifyParameterLogResponseBodyItems) *DescribeSearchModifyParameterLogResponseBody
	GetItems() []*DescribeSearchModifyParameterLogResponseBodyItems
	SetRequestId(v string) *DescribeSearchModifyParameterLogResponseBody
	GetRequestId() *string
}

type DescribeSearchModifyParameterLogResponseBody struct {
	// The type of the Search engine.
	//
	// example:
	//
	// polardb_search
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the Search DPI engine. This is not the MySQL version of the primary cluster.
	//
	// example:
	//
	// 3.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The parameter modification history. The records are returned in descending order of record IDs.
	Items []*DescribeSearchModifyParameterLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// e3a24663-eee5-4009-b635-2e9e8d052e74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSearchModifyParameterLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSearchModifyParameterLogResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSearchModifyParameterLogResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeSearchModifyParameterLogResponseBody) GetItems() []*DescribeSearchModifyParameterLogResponseBodyItems {
	return s.Items
}

func (s *DescribeSearchModifyParameterLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSearchModifyParameterLogResponseBody) SetEngine(v string) *DescribeSearchModifyParameterLogResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBody) SetEngineVersion(v string) *DescribeSearchModifyParameterLogResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBody) SetItems(v []*DescribeSearchModifyParameterLogResponseBodyItems) *DescribeSearchModifyParameterLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBody) SetRequestId(v string) *DescribeSearchModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBody) Validate() error {
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

type DescribeSearchModifyParameterLogResponseBodyItems struct {
	// The time when the parameter was modified. The time is displayed in the yyyy-MM-ddTHH:mm:ssZ format in UTC.
	//
	// example:
	//
	// 2026-10-14T07:04:29Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The new value of the parameter.
	//
	// example:
	//
	// 4
	NewParameterValue *string `json:"NewParameterValue,omitempty" xml:"NewParameterValue,omitempty"`
	// The original value of the parameter.
	//
	// example:
	//
	// 3
	OldParameterValue *string `json:"OldParameterValue,omitempty" xml:"OldParameterValue,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// http.compression_level
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The modification status. Valid values:
	//
	// - Syncing: The modification has not taken effect.
	//
	// - Applied: The modification has taken effect.
	//
	// example:
	//
	// Applied
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSearchModifyParameterLogResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchModifyParameterLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) GetNewParameterValue() *string {
	return s.NewParameterValue
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) GetOldParameterValue() *string {
	return s.OldParameterValue
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) SetModifyTime(v string) *DescribeSearchModifyParameterLogResponseBodyItems {
	s.ModifyTime = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) SetNewParameterValue(v string) *DescribeSearchModifyParameterLogResponseBodyItems {
	s.NewParameterValue = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) SetOldParameterValue(v string) *DescribeSearchModifyParameterLogResponseBodyItems {
	s.OldParameterValue = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) SetParameterName(v string) *DescribeSearchModifyParameterLogResponseBodyItems {
	s.ParameterName = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) SetStatus(v string) *DescribeSearchModifyParameterLogResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
