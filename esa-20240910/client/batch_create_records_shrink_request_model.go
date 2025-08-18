// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordListShrink(v string) *BatchCreateRecordsShrinkRequest
	GetRecordListShrink() *string
	SetSiteId(v int64) *BatchCreateRecordsShrinkRequest
	GetSiteId() *int64
}

type BatchCreateRecordsShrinkRequest struct {
	// The list of DNS records to be created.
	//
	// This parameter is required.
	RecordListShrink *string `json:"RecordList,omitempty" xml:"RecordList,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchCreateRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsShrinkRequest) GetRecordListShrink() *string {
	return s.RecordListShrink
}

func (s *BatchCreateRecordsShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchCreateRecordsShrinkRequest) SetRecordListShrink(v string) *BatchCreateRecordsShrinkRequest {
	s.RecordListShrink = &v
	return s
}

func (s *BatchCreateRecordsShrinkRequest) SetSiteId(v int64) *BatchCreateRecordsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
