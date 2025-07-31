// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupBucketsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetBackupBucketsListResponseBodyData) *GetBackupBucketsListResponseBody
	GetData() []*GetBackupBucketsListResponseBodyData
	SetRequestId(v string) *GetBackupBucketsListResponseBody
	GetRequestId() *string
}

type GetBackupBucketsListResponseBody struct {
	Data []*GetBackupBucketsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupBucketsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupBucketsListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListResponseBody) GetData() []*GetBackupBucketsListResponseBodyData {
	return s.Data
}

func (s *GetBackupBucketsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupBucketsListResponseBody) SetData(v []*GetBackupBucketsListResponseBodyData) *GetBackupBucketsListResponseBody {
	s.Data = v
	return s
}

func (s *GetBackupBucketsListResponseBody) SetRequestId(v string) *GetBackupBucketsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupBucketsListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBackupBucketsListResponseBodyData struct {
	// example:
	//
	// gj-bucket1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetBackupBucketsListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBackupBucketsListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetBackupBucketsListResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetBackupBucketsListResponseBodyData) SetBucket(v string) *GetBackupBucketsListResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetBackupBucketsListResponseBodyData) SetRegion(v string) *GetBackupBucketsListResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetBackupBucketsListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
