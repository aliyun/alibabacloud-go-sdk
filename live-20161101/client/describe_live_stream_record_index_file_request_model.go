// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordIndexFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamRecordIndexFileRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamRecordIndexFileRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamRecordIndexFileRequest
	GetOwnerId() *int64
	SetRecordId(v string) *DescribeLiveStreamRecordIndexFileRequest
	GetRecordId() *string
	SetSecurityToken(v string) *DescribeLiveStreamRecordIndexFileRequest
	GetSecurityToken() *string
	SetStreamName(v string) *DescribeLiveStreamRecordIndexFileRequest
	GetStreamName() *string
}

type DescribeLiveStreamRecordIndexFileRequest struct {
	// System-defined parameter. Value: **DescribeLiveStreamRecordIndexFile**.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// ## [](#)Usage notes
	//
	// ApsaraVideo Live stores the configuration information of an M3U8 index file for six months. You can query only index files created in the previous six months. M3U8 index files are stored in Object Storage Service (OSS) buckets. The retention period is determined by the storage configuration of the OSS buckets.
	//
	// ## [](#qps-)QPS limit
	//
	// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the index file.
	//
	// >  You can call the [DescribeLiveStreamRecordIndexFiles](https://help.aliyun.com/document_detail/2847890.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c4d7f0a4-b506-43f9-8de3-07732c3f****
	RecordId      *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The main domain of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamRecordIndexFileRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamRecordIndexFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamRecordIndexFileRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeLiveStreamRecordIndexFileRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamRecordIndexFileRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetAppName(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetDomainName(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetOwnerId(v int64) *DescribeLiveStreamRecordIndexFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetRecordId(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetSecurityToken(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetStreamName(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) Validate() error {
	return dara.Validate(s)
}
