// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskAttachments(v []*DescribeDiskAttachmentsResponseBodyDiskAttachments) *DescribeDiskAttachmentsResponseBody
	GetDiskAttachments() []*DescribeDiskAttachmentsResponseBodyDiskAttachments
	SetMaxResults(v int32) *DescribeDiskAttachmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDiskAttachmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDiskAttachmentsResponseBody
	GetRequestId() *string
}

type DescribeDiskAttachmentsResponseBody struct {
	// The list of disk attachments.
	DiskAttachments []*DescribeDiskAttachmentsResponseBodyDiskAttachments `json:"DiskAttachments,omitempty" xml:"DiskAttachments,omitempty" type:"Repeated"`
	// The maximum number of attachments returned in a single request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that all attachments have been returned.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskAttachmentsResponseBody) GetDiskAttachments() []*DescribeDiskAttachmentsResponseBodyDiskAttachments {
	return s.DiskAttachments
}

func (s *DescribeDiskAttachmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiskAttachmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskAttachmentsResponseBody) SetDiskAttachments(v []*DescribeDiskAttachmentsResponseBodyDiskAttachments) *DescribeDiskAttachmentsResponseBody {
	s.DiskAttachments = v
	return s
}

func (s *DescribeDiskAttachmentsResponseBody) SetMaxResults(v int32) *DescribeDiskAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBody) SetNextToken(v string) *DescribeDiskAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBody) SetRequestId(v string) *DescribeDiskAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBody) Validate() error {
	if s.DiskAttachments != nil {
		for _, item := range s.DiskAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskAttachmentsResponseBodyDiskAttachments struct {
	// The time when the disk was successfully attached, in ISO 8601 format (UTC).
	//
	// example:
	//
	// 2026-10-12T06:08:54Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The device name on the instance.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the instance to which the disk is attached.
	//
	// example:
	//
	// i-bp18pnlg1ds9rky4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDiskAttachmentsResponseBodyDiskAttachments) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskAttachmentsResponseBodyDiskAttachments) GoString() string {
	return s.String()
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) GetDevice() *string {
	return s.Device
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) SetAttachedTime(v string) *DescribeDiskAttachmentsResponseBodyDiskAttachments {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) SetDevice(v string) *DescribeDiskAttachmentsResponseBodyDiskAttachments {
	s.Device = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) SetDiskId(v string) *DescribeDiskAttachmentsResponseBodyDiskAttachments {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) SetInstanceId(v string) *DescribeDiskAttachmentsResponseBodyDiskAttachments {
	s.InstanceId = &v
	return s
}

func (s *DescribeDiskAttachmentsResponseBodyDiskAttachments) Validate() error {
	return dara.Validate(s)
}
