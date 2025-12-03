// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetInfo(v *DescribeDatasetInfoResponseBodyDatasetInfo) *DescribeDatasetInfoResponseBody
	GetDatasetInfo() *DescribeDatasetInfoResponseBodyDatasetInfo
	SetRequestId(v string) *DescribeDatasetInfoResponseBody
	GetRequestId() *string
}

type DescribeDatasetInfoResponseBody struct {
	// The dataset info.
	DatasetInfo *DescribeDatasetInfoResponseBodyDatasetInfo `json:"DatasetInfo,omitempty" xml:"DatasetInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A2C8F75E-EE84-5C64-960F-45C8********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDatasetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatasetInfoResponseBody) GetDatasetInfo() *DescribeDatasetInfoResponseBodyDatasetInfo {
	return s.DatasetInfo
}

func (s *DescribeDatasetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatasetInfoResponseBody) SetDatasetInfo(v *DescribeDatasetInfoResponseBodyDatasetInfo) *DescribeDatasetInfoResponseBody {
	s.DatasetInfo = v
	return s
}

func (s *DescribeDatasetInfoResponseBody) SetRequestId(v string) *DescribeDatasetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatasetInfoResponseBody) Validate() error {
	if s.DatasetInfo != nil {
		if err := s.DatasetInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDatasetInfoResponseBodyDatasetInfo struct {
	// The creation time (UTC) of the dataset.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the dataset.
	//
	// example:
	//
	// 62b91a790a693238********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The type of the dataset. Valid values:
	//
	// 	- JWT_BLOCKING: a JSON Web Token (JWT) blacklist
	//
	// 	- IP_WHITELIST_CIDR : an IP address whitelist
	//
	// 	- PARAMETER_ACCESS : parameter-based access control
	//
	// example:
	//
	// JWT_BLOCKING
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last modification time (UTC) of the dataset.
	//
	// example:
	//
	// 2022-09-21T12:58:43Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeDatasetInfoResponseBodyDatasetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetInfoResponseBodyDatasetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) GetDatasetType() *string {
	return s.DatasetType
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) SetCreatedTime(v string) *DescribeDatasetInfoResponseBodyDatasetInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) SetDatasetId(v string) *DescribeDatasetInfoResponseBodyDatasetInfo {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) SetDatasetName(v string) *DescribeDatasetInfoResponseBodyDatasetInfo {
	s.DatasetName = &v
	return s
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) SetDatasetType(v string) *DescribeDatasetInfoResponseBodyDatasetInfo {
	s.DatasetType = &v
	return s
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) SetDescription(v string) *DescribeDatasetInfoResponseBodyDatasetInfo {
	s.Description = &v
	return s
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) SetModifiedTime(v string) *DescribeDatasetInfoResponseBodyDatasetInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeDatasetInfoResponseBodyDatasetInfo) Validate() error {
	return dara.Validate(s)
}
