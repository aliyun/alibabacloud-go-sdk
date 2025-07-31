// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetBucketsListResponseBodyData) *GetBucketsListResponseBody
	GetData() []*GetBucketsListResponseBodyData
	SetRequestId(v string) *GetBucketsListResponseBody
	GetRequestId() *string
}

type GetBucketsListResponseBody struct {
	Data []*GetBucketsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBucketsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketsListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketsListResponseBody) GetData() []*GetBucketsListResponseBodyData {
	return s.Data
}

func (s *GetBucketsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBucketsListResponseBody) SetData(v []*GetBucketsListResponseBodyData) *GetBucketsListResponseBody {
	s.Data = v
	return s
}

func (s *GetBucketsListResponseBody) SetRequestId(v string) *GetBucketsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBucketsListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBucketsListResponseBodyData struct {
	// example:
	//
	// bucket_test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetBucketsListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBucketsListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBucketsListResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetBucketsListResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetBucketsListResponseBodyData) SetBucket(v string) *GetBucketsListResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetBucketsListResponseBodyData) SetRegion(v string) *GetBucketsListResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetBucketsListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
