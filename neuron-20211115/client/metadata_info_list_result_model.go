// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadataInfoListResult interface {
	dara.Model
	String() string
	GoString() string
	SetMetadataInfos(v []*MetadataInfo) *MetadataInfoListResult
	GetMetadataInfos() []*MetadataInfo
	SetRequestId(v string) *MetadataInfoListResult
	GetRequestId() *string
	SetTotal(v int64) *MetadataInfoListResult
	GetTotal() *int64
}

type MetadataInfoListResult struct {
	MetadataInfos []*MetadataInfo `json:"metadataInfos,omitempty" xml:"metadataInfos,omitempty" type:"Repeated"`
	RequestId     *string         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MetadataInfoListResult) String() string {
	return dara.Prettify(s)
}

func (s MetadataInfoListResult) GoString() string {
	return s.String()
}

func (s *MetadataInfoListResult) GetMetadataInfos() []*MetadataInfo {
	return s.MetadataInfos
}

func (s *MetadataInfoListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MetadataInfoListResult) GetTotal() *int64 {
	return s.Total
}

func (s *MetadataInfoListResult) SetMetadataInfos(v []*MetadataInfo) *MetadataInfoListResult {
	s.MetadataInfos = v
	return s
}

func (s *MetadataInfoListResult) SetRequestId(v string) *MetadataInfoListResult {
	s.RequestId = &v
	return s
}

func (s *MetadataInfoListResult) SetTotal(v int64) *MetadataInfoListResult {
	s.Total = &v
	return s
}

func (s *MetadataInfoListResult) Validate() error {
	if s.MetadataInfos != nil {
		for _, item := range s.MetadataInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
