// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetInfoPublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*ListAssetInfoPublishResponseBodyAssetList) *ListAssetInfoPublishResponseBody
	GetAssetList() []*ListAssetInfoPublishResponseBodyAssetList
	SetRequestId(v string) *ListAssetInfoPublishResponseBody
	GetRequestId() *string
}

type ListAssetInfoPublishResponseBody struct {
	// The server list information.
	AssetList []*ListAssetInfoPublishResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB483977
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssetInfoPublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetInfoPublishResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetInfoPublishResponseBody) GetAssetList() []*ListAssetInfoPublishResponseBodyAssetList {
	return s.AssetList
}

func (s *ListAssetInfoPublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetInfoPublishResponseBody) SetAssetList(v []*ListAssetInfoPublishResponseBodyAssetList) *ListAssetInfoPublishResponseBody {
	s.AssetList = v
	return s
}

func (s *ListAssetInfoPublishResponseBody) SetRequestId(v string) *ListAssetInfoPublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetInfoPublishResponseBody) Validate() error {
	if s.AssetList != nil {
		for _, item := range s.AssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssetInfoPublishResponseBodyAssetList struct {
	// The current client version.
	//
	// example:
	//
	// 0.0.8
	CurVersion *string `json:"CurVersion,omitempty" xml:"CurVersion,omitempty"`
	// The time of the last upgrade. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1724522400000
	LastUpgradeTime *int64 `json:"LastUpgradeTime,omitempty" xml:"LastUpgradeTime,omitempty"`
	// The client release status. Valid values:
	//
	// - **0**: Not started.
	//
	// - **1**: Publishing.
	//
	// - **2**: Publishing completed.
	//
	// - **3**: Publishing paused.
	//
	// - **4**: Force upgrading.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the upgrade is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	UpgradeEnable *bool `json:"UpgradeEnable,omitempty" xml:"UpgradeEnable,omitempty"`
	// The UUID of the Security Center asset.
	//
	// example:
	//
	// 2a98f149-0256-414c-a29a-a69f8a75****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAssetInfoPublishResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s ListAssetInfoPublishResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *ListAssetInfoPublishResponseBodyAssetList) GetCurVersion() *string {
	return s.CurVersion
}

func (s *ListAssetInfoPublishResponseBodyAssetList) GetLastUpgradeTime() *int64 {
	return s.LastUpgradeTime
}

func (s *ListAssetInfoPublishResponseBodyAssetList) GetStatus() *int32 {
	return s.Status
}

func (s *ListAssetInfoPublishResponseBodyAssetList) GetUpgradeEnable() *bool {
	return s.UpgradeEnable
}

func (s *ListAssetInfoPublishResponseBodyAssetList) GetUuid() *string {
	return s.Uuid
}

func (s *ListAssetInfoPublishResponseBodyAssetList) SetCurVersion(v string) *ListAssetInfoPublishResponseBodyAssetList {
	s.CurVersion = &v
	return s
}

func (s *ListAssetInfoPublishResponseBodyAssetList) SetLastUpgradeTime(v int64) *ListAssetInfoPublishResponseBodyAssetList {
	s.LastUpgradeTime = &v
	return s
}

func (s *ListAssetInfoPublishResponseBodyAssetList) SetStatus(v int32) *ListAssetInfoPublishResponseBodyAssetList {
	s.Status = &v
	return s
}

func (s *ListAssetInfoPublishResponseBodyAssetList) SetUpgradeEnable(v bool) *ListAssetInfoPublishResponseBodyAssetList {
	s.UpgradeEnable = &v
	return s
}

func (s *ListAssetInfoPublishResponseBodyAssetList) SetUuid(v string) *ListAssetInfoPublishResponseBodyAssetList {
	s.Uuid = &v
	return s
}

func (s *ListAssetInfoPublishResponseBodyAssetList) Validate() error {
	return dara.Validate(s)
}
