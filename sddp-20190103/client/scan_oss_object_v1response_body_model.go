// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanOssObjectV1ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ScanOssObjectV1ResponseBody
	GetId() *int64
	SetRequestId(v string) *ScanOssObjectV1ResponseBody
	GetRequestId() *string
}

type ScanOssObjectV1ResponseBody struct {
	// The ID of the identification task that is returned after the identification task is created.
	//
	// example:
	//
	// 268
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScanOssObjectV1ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanOssObjectV1ResponseBody) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1ResponseBody) GetId() *int64 {
	return s.Id
}

func (s *ScanOssObjectV1ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanOssObjectV1ResponseBody) SetId(v int64) *ScanOssObjectV1ResponseBody {
	s.Id = &v
	return s
}

func (s *ScanOssObjectV1ResponseBody) SetRequestId(v string) *ScanOssObjectV1ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanOssObjectV1ResponseBody) Validate() error {
	return dara.Validate(s)
}
