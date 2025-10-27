// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvice(v string) *DescribeCheckWarningDetailResponseBody
	GetAdvice() *string
	SetCheckDetailAssetInfo(v []map[string]*string) *DescribeCheckWarningDetailResponseBody
	GetCheckDetailAssetInfo() []map[string]*string
	SetCheckDetailColumns(v []*DescribeCheckWarningDetailResponseBodyCheckDetailColumns) *DescribeCheckWarningDetailResponseBody
	GetCheckDetailColumns() []*DescribeCheckWarningDetailResponseBodyCheckDetailColumns
	SetCheckId(v int64) *DescribeCheckWarningDetailResponseBody
	GetCheckId() *int64
	SetDescription(v string) *DescribeCheckWarningDetailResponseBody
	GetDescription() *string
	SetItem(v string) *DescribeCheckWarningDetailResponseBody
	GetItem() *string
	SetLevel(v string) *DescribeCheckWarningDetailResponseBody
	GetLevel() *string
	SetPrompt(v string) *DescribeCheckWarningDetailResponseBody
	GetPrompt() *string
	SetRequestId(v string) *DescribeCheckWarningDetailResponseBody
	GetRequestId() *string
	SetType(v string) *DescribeCheckWarningDetailResponseBody
	GetType() *string
}

type DescribeCheckWarningDetailResponseBody struct {
	// The suggestion for the management of the risk item.
	//
	// example:
	//
	// You can fix it in the following ways:↵1. To configure authentication for redis service, click the redis.conf Configure complex password in requirepass, and then restart redis.↵2. In redis configuration file redis.conf The configuration is as follows: bind 127.0.0.1, only allow local access, and then restart redis
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// List of asset details to check.
	CheckDetailAssetInfo []map[string]*string `json:"CheckDetailAssetInfo,omitempty" xml:"CheckDetailAssetInfo,omitempty" type:"Repeated"`
	// Detection content details.
	CheckDetailColumns []*DescribeCheckWarningDetailResponseBodyCheckDetailColumns `json:"CheckDetailColumns,omitempty" xml:"CheckDetailColumns,omitempty" type:"Repeated"`
	// The ID of the check item.
	//
	// example:
	//
	// 946
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The additional information about the risk item.
	//
	// example:
	//
	// The redis port is open to the outside world and there is no authentication option configured. In addition to directly obtaining all the information in the database, unauthorized users can also attack the system through unauthorized access vulnerability.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// Redis unauthorized access
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// 	- **high**: The item is a high-risk item and is highlighted in red.
	//
	// 	- **medium**: The item is a medium-risk item and is highlighted in orange.
	//
	// 	- **low**: The item is a low-risk item and is highlighted in gray.
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The prompt for the risk item.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the check item. Valid values:
	//
	// 	- **hc_exploit**: unauthorized access
	//
	// 	- **hc_djbh**: classified protection compliance
	//
	// 	- **hc_best_secruity**: best security practice
	//
	// 	- **hc_container**: container security
	//
	// 	- **hc_custom**: custom baseline
	//
	// 	- **cis**: Center for Internet Security (CIS) compliance
	//
	// 	- **weak_password**: weak password
	//
	// example:
	//
	// hc_exploit
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCheckWarningDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponseBody) GetAdvice() *string {
	return s.Advice
}

func (s *DescribeCheckWarningDetailResponseBody) GetCheckDetailAssetInfo() []map[string]*string {
	return s.CheckDetailAssetInfo
}

func (s *DescribeCheckWarningDetailResponseBody) GetCheckDetailColumns() []*DescribeCheckWarningDetailResponseBodyCheckDetailColumns {
	return s.CheckDetailColumns
}

func (s *DescribeCheckWarningDetailResponseBody) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckWarningDetailResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeCheckWarningDetailResponseBody) GetItem() *string {
	return s.Item
}

func (s *DescribeCheckWarningDetailResponseBody) GetLevel() *string {
	return s.Level
}

func (s *DescribeCheckWarningDetailResponseBody) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeCheckWarningDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckWarningDetailResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeCheckWarningDetailResponseBody) SetAdvice(v string) *DescribeCheckWarningDetailResponseBody {
	s.Advice = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetCheckDetailAssetInfo(v []map[string]*string) *DescribeCheckWarningDetailResponseBody {
	s.CheckDetailAssetInfo = v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetCheckDetailColumns(v []*DescribeCheckWarningDetailResponseBodyCheckDetailColumns) *DescribeCheckWarningDetailResponseBody {
	s.CheckDetailColumns = v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetCheckId(v int64) *DescribeCheckWarningDetailResponseBody {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetDescription(v string) *DescribeCheckWarningDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetItem(v string) *DescribeCheckWarningDetailResponseBody {
	s.Item = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetLevel(v string) *DescribeCheckWarningDetailResponseBody {
	s.Level = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetPrompt(v string) *DescribeCheckWarningDetailResponseBody {
	s.Prompt = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetRequestId(v string) *DescribeCheckWarningDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetType(v string) *DescribeCheckWarningDetailResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) Validate() error {
	if s.CheckDetailColumns != nil {
		for _, item := range s.CheckDetailColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCheckWarningDetailResponseBodyCheckDetailColumns struct {
	// Detection content list.
	Grids []*DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids `json:"Grids,omitempty" xml:"Grids,omitempty" type:"Repeated"`
	// Key to detect content.
	//
	// example:
	//
	// Containername
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The detection content key corresponds to the display name.
	//
	// example:
	//
	// ContainerName
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// Display type. Value:
	//
	// - **grid**: Detection grid
	//
	// - **text**: text
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCheckWarningDetailResponseBodyCheckDetailColumns) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningDetailResponseBodyCheckDetailColumns) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) GetGrids() []*DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids {
	return s.Grids
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) GetKey() *string {
	return s.Key
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) GetShowName() *string {
	return s.ShowName
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) GetType() *string {
	return s.Type
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) SetGrids(v []*DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) *DescribeCheckWarningDetailResponseBodyCheckDetailColumns {
	s.Grids = v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) SetKey(v string) *DescribeCheckWarningDetailResponseBodyCheckDetailColumns {
	s.Key = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) SetShowName(v string) *DescribeCheckWarningDetailResponseBodyCheckDetailColumns {
	s.ShowName = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) SetType(v string) *DescribeCheckWarningDetailResponseBodyCheckDetailColumns {
	s.Type = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumns) Validate() error {
	if s.Grids != nil {
		for _, item := range s.Grids {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids struct {
	// Key to detect content.
	//
	// example:
	//
	// Username
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The detection content key corresponds to the display name.
	//
	// example:
	//
	// UserName
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// Display type. Value:
	//
	// - **grid**: Detection grid
	//
	// - **text**: text
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) GetKey() *string {
	return s.Key
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) GetShowName() *string {
	return s.ShowName
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) GetType() *string {
	return s.Type
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) SetKey(v string) *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids {
	s.Key = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) SetShowName(v string) *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids {
	s.ShowName = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) SetType(v string) *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids {
	s.Type = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBodyCheckDetailColumnsGrids) Validate() error {
	return dara.Validate(s)
}
