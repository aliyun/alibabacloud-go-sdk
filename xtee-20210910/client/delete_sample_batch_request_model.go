// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteSampleBatchRequest
	GetLang() *string
	SetIds(v string) *DeleteSampleBatchRequest
	GetIds() *string
	SetRegId(v string) *DeleteSampleBatchRequest
	GetRegId() *string
	SetVersions(v string) *DeleteSampleBatchRequest
	GetVersions() *string
}

type DeleteSampleBatchRequest struct {
	// The language type for the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IDs of the samples to delete in batches.
	//
	// example:
	//
	// 324,343
	Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The list of versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,1
	Versions *string `json:"versions,omitempty" xml:"versions,omitempty"`
}

func (s DeleteSampleBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteSampleBatchRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSampleBatchRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteSampleBatchRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteSampleBatchRequest) GetVersions() *string {
	return s.Versions
}

func (s *DeleteSampleBatchRequest) SetLang(v string) *DeleteSampleBatchRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSampleBatchRequest) SetIds(v string) *DeleteSampleBatchRequest {
	s.Ids = &v
	return s
}

func (s *DeleteSampleBatchRequest) SetRegId(v string) *DeleteSampleBatchRequest {
	s.RegId = &v
	return s
}

func (s *DeleteSampleBatchRequest) SetVersions(v string) *DeleteSampleBatchRequest {
	s.Versions = &v
	return s
}

func (s *DeleteSampleBatchRequest) Validate() error {
	return dara.Validate(s)
}
