// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentAssets(v []*DescribeComponentAssetsResponseBodyComponentAssets) *DescribeComponentAssetsResponseBody
	GetComponentAssets() []*DescribeComponentAssetsResponseBodyComponentAssets
	SetRequestId(v string) *DescribeComponentAssetsResponseBody
	GetRequestId() *string
}

type DescribeComponentAssetsResponseBody struct {
	// The information about the assets.
	ComponentAssets []*DescribeComponentAssetsResponseBodyComponentAssets `json:"ComponentAssets,omitempty" xml:"ComponentAssets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BFEFB76D-DD0E-5529-BD57-0DAC10B9B30F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponseBody) GetComponentAssets() []*DescribeComponentAssetsResponseBodyComponentAssets {
	return s.ComponentAssets
}

func (s *DescribeComponentAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentAssetsResponseBody) SetComponentAssets(v []*DescribeComponentAssetsResponseBodyComponentAssets) *DescribeComponentAssetsResponseBody {
	s.ComponentAssets = v
	return s
}

func (s *DescribeComponentAssetsResponseBody) SetRequestId(v string) *DescribeComponentAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentAssetsResponseBody) Validate() error {
	if s.ComponentAssets != nil {
		for _, item := range s.ComponentAssets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComponentAssetsResponseBodyComponentAssets struct {
	// The UUID of the asset.
	//
	// example:
	//
	// ff6fe161-93e2-464c-a326-fxxxxxx
	AssetUuid *string `json:"AssetUuid,omitempty" xml:"AssetUuid,omitempty"`
	// The name of the component to which the asset belongs.
	//
	// example:
	//
	// pyhton3
	Componentname *string `json:"Componentname,omitempty" xml:"Componentname,omitempty"`
	// The time when the asset was created. The time is in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2023-03-23T14:38Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the asset was modified. The time is in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2023-03-23T14:38Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 7xx
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// test asset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the asset in the JSON string format. DescribeComponentAssetForm
	//
	// >  For more information, see [DescribeComponentAssetForm](~~DescribeComponentAssetForm~~).
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "authMethod",
	//
	//         "value": "ak"
	//
	//     },
	//
	//     {
	//
	//         "name": "accessKeyId",
	//
	//         "value": "xxxxxxx"
	//
	//     },
	//
	//     {
	//
	//         "name": "accessKeySecret",
	//
	//         "value": "xxxxx"
	//
	//     },
	//
	//     {
	//
	//         "name": "roleArn",
	//
	//         "value": ""
	//
	//     }
	//
	// ]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s DescribeComponentAssetsResponseBodyComponentAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetsResponseBodyComponentAssets) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetAssetUuid() *string {
	return s.AssetUuid
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetComponentname() *string {
	return s.Componentname
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetId() *int64 {
	return s.Id
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetName() *string {
	return s.Name
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) GetParams() *string {
	return s.Params
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetAssetUuid(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.AssetUuid = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetComponentname(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Componentname = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetGmtCreate(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.GmtCreate = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetGmtModified(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.GmtModified = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetId(v int64) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Id = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetName(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Name = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetParams(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Params = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) Validate() error {
	return dara.Validate(s)
}
