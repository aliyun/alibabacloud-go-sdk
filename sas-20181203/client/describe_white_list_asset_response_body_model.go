// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*DescribeWhiteListAssetResponseBodyAssets) *DescribeWhiteListAssetResponseBody
	GetAssets() []*DescribeWhiteListAssetResponseBodyAssets
	SetRequestId(v string) *DescribeWhiteListAssetResponseBody
	GetRequestId() *string
}

type DescribeWhiteListAssetResponseBody struct {
	// The information about the servers.
	Assets []*DescribeWhiteListAssetResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C61****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWhiteListAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAssetResponseBody) GetAssets() []*DescribeWhiteListAssetResponseBodyAssets {
	return s.Assets
}

func (s *DescribeWhiteListAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListAssetResponseBody) SetAssets(v []*DescribeWhiteListAssetResponseBodyAssets) *DescribeWhiteListAssetResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeWhiteListAssetResponseBody) SetRequestId(v string) *DescribeWhiteListAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBody) Validate() error {
	if s.Assets != nil {
		for _, item := range s.Assets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWhiteListAssetResponseBodyAssets struct {
	// Indicates whether the server can be selected. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllowSelected *int32 `json:"AllowSelected,omitempty" xml:"AllowSelected,omitempty"`
	// The group ID of the server.
	//
	// example:
	//
	// 11028551
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 121.41.XX.XX
	MachineIp *string `json:"MachineIp,omitempty" xml:"MachineIp,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// test
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// Indicates whether the server is selected. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	Selected *int32 `json:"Selected,omitempty" xml:"Selected,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 5c081b02-f66a-47a4-bd2f-79ee3eaf806a
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWhiteListAssetResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAssetResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetAllowSelected() *int32 {
	return s.AllowSelected
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetId() *int64 {
	return s.Id
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetMachineIp() *string {
	return s.MachineIp
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetMachineName() *string {
	return s.MachineName
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetSelected() *int32 {
	return s.Selected
}

func (s *DescribeWhiteListAssetResponseBodyAssets) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetAllowSelected(v int32) *DescribeWhiteListAssetResponseBodyAssets {
	s.AllowSelected = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetGroupId(v int64) *DescribeWhiteListAssetResponseBodyAssets {
	s.GroupId = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetId(v int64) *DescribeWhiteListAssetResponseBodyAssets {
	s.Id = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetMachineIp(v string) *DescribeWhiteListAssetResponseBodyAssets {
	s.MachineIp = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetMachineName(v string) *DescribeWhiteListAssetResponseBodyAssets {
	s.MachineName = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetSelected(v int32) *DescribeWhiteListAssetResponseBodyAssets {
	s.Selected = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) SetUuid(v string) *DescribeWhiteListAssetResponseBodyAssets {
	s.Uuid = &v
	return s
}

func (s *DescribeWhiteListAssetResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}
