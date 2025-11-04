// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateVodPackagingAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *BatchCreateVodPackagingAssetResponseBody
	GetGroupName() *string
	SetRequestId(v string) *BatchCreateVodPackagingAssetResponseBody
	GetRequestId() *string
	SetResultList(v []*BatchCreateVodPackagingAssetResponseBodyResultList) *BatchCreateVodPackagingAssetResponseBody
	GetResultList() []*BatchCreateVodPackagingAssetResponseBodyResultList
}

type BatchCreateVodPackagingAssetResponseBody struct {
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of asset ingestion.
	ResultList []*BatchCreateVodPackagingAssetResponseBodyResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s BatchCreateVodPackagingAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *BatchCreateVodPackagingAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateVodPackagingAssetResponseBody) GetResultList() []*BatchCreateVodPackagingAssetResponseBodyResultList {
	return s.ResultList
}

func (s *BatchCreateVodPackagingAssetResponseBody) SetGroupName(v string) *BatchCreateVodPackagingAssetResponseBody {
	s.GroupName = &v
	return s
}

func (s *BatchCreateVodPackagingAssetResponseBody) SetRequestId(v string) *BatchCreateVodPackagingAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateVodPackagingAssetResponseBody) SetResultList(v []*BatchCreateVodPackagingAssetResponseBodyResultList) *BatchCreateVodPackagingAssetResponseBody {
	s.ResultList = v
	return s
}

func (s *BatchCreateVodPackagingAssetResponseBody) Validate() error {
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateVodPackagingAssetResponseBodyResultList struct {
	// The information about the ingested asset.
	Asset *VodPackagingAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// The error code for failed ingestion.
	//
	// example:
	//
	// InvalidParameter.PackagingAssetAlreadyExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message for failed ingestion.
	//
	// example:
	//
	// The specified packagingAsset "inputMovie" already exists
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s BatchCreateVodPackagingAssetResponseBodyResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) GetAsset() *VodPackagingAsset {
	return s.Asset
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) GetCode() *string {
	return s.Code
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) SetAsset(v *VodPackagingAsset) *BatchCreateVodPackagingAssetResponseBodyResultList {
	s.Asset = v
	return s
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) SetCode(v string) *BatchCreateVodPackagingAssetResponseBodyResultList {
	s.Code = &v
	return s
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) SetMessage(v string) *BatchCreateVodPackagingAssetResponseBodyResultList {
	s.Message = &v
	return s
}

func (s *BatchCreateVodPackagingAssetResponseBodyResultList) Validate() error {
	if s.Asset != nil {
		if err := s.Asset.Validate(); err != nil {
			return err
		}
	}
	return nil
}
