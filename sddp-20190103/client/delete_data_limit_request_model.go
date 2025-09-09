// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DeleteDataLimitRequest
	GetFeatureType() *int32
	SetId(v int64) *DeleteDataLimitRequest
	GetId() *int64
	SetLang(v string) *DeleteDataLimitRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteDataLimitRequest
	GetSourceIp() *string
}

type DeleteDataLimitRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The ID of the data asset.
	//
	// You can call the DescribeDataLimits operation to query the IDs of data assets. The value of the Id response parameter indicates the ID of a data asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteDataLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLimitRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DeleteDataLimitRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataLimitRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteDataLimitRequest) SetFeatureType(v int32) *DeleteDataLimitRequest {
	s.FeatureType = &v
	return s
}

func (s *DeleteDataLimitRequest) SetId(v int64) *DeleteDataLimitRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataLimitRequest) SetLang(v string) *DeleteDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataLimitRequest) SetSourceIp(v string) *DeleteDataLimitRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteDataLimitRequest) Validate() error {
	return dara.Validate(s)
}
