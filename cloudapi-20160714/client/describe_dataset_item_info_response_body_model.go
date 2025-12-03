// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetItemInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetItemInfo(v *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) *DescribeDatasetItemInfoResponseBody
	GetDatasetItemInfo() *DescribeDatasetItemInfoResponseBodyDatasetItemInfo
	SetRequestId(v string) *DescribeDatasetItemInfoResponseBody
	GetRequestId() *string
}

type DescribeDatasetItemInfoResponseBody struct {
	// The Dataset information.
	DatasetItemInfo *DescribeDatasetItemInfoResponseBodyDatasetItemInfo `json:"DatasetItemInfo,omitempty" xml:"DatasetItemInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5BAFA85F-38E3-5D9E-9E32-4B09********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDatasetItemInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemInfoResponseBody) GetDatasetItemInfo() *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	return s.DatasetItemInfo
}

func (s *DescribeDatasetItemInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatasetItemInfoResponseBody) SetDatasetItemInfo(v *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) *DescribeDatasetItemInfoResponseBody {
	s.DatasetItemInfo = v
	return s
}

func (s *DescribeDatasetItemInfoResponseBody) SetRequestId(v string) *DescribeDatasetItemInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBody) Validate() error {
	if s.DatasetItemInfo != nil {
		if err := s.DatasetItemInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDatasetItemInfoResponseBodyDatasetItemInfo struct {
	// The creation time (UTC) of the data entry.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the dataset.
	//
	// example:
	//
	// 626238665db4a5140eea3e40********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The ID of the data entry.
	//
	// example:
	//
	// 5045****
	DatasetItemId *string `json:"DatasetItemId,omitempty" xml:"DatasetItemId,omitempty"`
	// The description of the data entry.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time in UTC when the data entry expires. The time is in the **yyyy-MM-ddTHH:mm:ssZ*	- format. If this parameter is empty, the data entry does not expire.
	//
	// example:
	//
	// 2022-09-22T12:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The last modification time (UTC) of the data entry.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The value of the data entry.
	//
	// example:
	//
	// 106.43.XXX.XXX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDatasetItemInfoResponseBodyDatasetItemInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetDatasetItemId() *string {
	return s.DatasetItemId
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetCreatedTime(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetDatasetId(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetDatasetItemId(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.DatasetItemId = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetDescription(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.Description = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetExpiredTime(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetModifiedTime(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) SetValue(v string) *DescribeDatasetItemInfoResponseBodyDatasetItemInfo {
	s.Value = &v
	return s
}

func (s *DescribeDatasetItemInfoResponseBodyDatasetItemInfo) Validate() error {
	return dara.Validate(s)
}
