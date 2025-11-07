// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListFaceVerifyInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFaceVerifyInfos(v []*DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) *DescribeListFaceVerifyInfosResponseBody
	GetFaceVerifyInfos() []*DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos
	SetItemsPerPage(v int32) *DescribeListFaceVerifyInfosResponseBody
	GetItemsPerPage() *int32
	SetPageNumber(v int32) *DescribeListFaceVerifyInfosResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeListFaceVerifyInfosResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeListFaceVerifyInfosResponseBody
	GetTotalCount() *int32
	SetTotalPages(v int32) *DescribeListFaceVerifyInfosResponseBody
	GetTotalPages() *int32
}

type DescribeListFaceVerifyInfosResponseBody struct {
	// List of face verification records.
	FaceVerifyInfos []*DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos `json:"FaceVerifyInfos,omitempty" xml:"FaceVerifyInfos,omitempty" type:"Repeated"`
	// Number of items per page.
	//
	// example:
	//
	// 20
	ItemsPerPage *int32 `json:"ItemsPerPage,omitempty" xml:"ItemsPerPage,omitempty"`
	// Pagination parameter: current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// ID of the current request.
	//
	// example:
	//
	// 1CC27D8E-24BF-5056-B14E-9F26719C9A8D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of verifications.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeListFaceVerifyInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyInfosResponseBody) GetFaceVerifyInfos() []*DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	return s.FaceVerifyInfos
}

func (s *DescribeListFaceVerifyInfosResponseBody) GetItemsPerPage() *int32 {
	return s.ItemsPerPage
}

func (s *DescribeListFaceVerifyInfosResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeListFaceVerifyInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListFaceVerifyInfosResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeListFaceVerifyInfosResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeListFaceVerifyInfosResponseBody) SetFaceVerifyInfos(v []*DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) *DescribeListFaceVerifyInfosResponseBody {
	s.FaceVerifyInfos = v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBody) SetItemsPerPage(v int32) *DescribeListFaceVerifyInfosResponseBody {
	s.ItemsPerPage = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBody) SetPageNumber(v int32) *DescribeListFaceVerifyInfosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBody) SetRequestId(v string) *DescribeListFaceVerifyInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBody) SetTotalCount(v int32) *DescribeListFaceVerifyInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBody) SetTotalPages(v int32) *DescribeListFaceVerifyInfosResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBody) Validate() error {
	if s.FaceVerifyInfos != nil {
		for _, item := range s.FaceVerifyInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos struct {
	// Business code.
	//
	// example:
	//
	// CLOUD_FACE
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Name.
	//
	// example:
	//
	// 赵四
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// ID number.
	//
	// example:
	//
	// 500382199805086199
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// ID of the certificate.
	//
	// example:
	//
	// sha8ff58e964152c4c4d21005fb98ecb
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Creation time of the face recognition record.
	//
	// example:
	//
	// 2022-10-02T11:16:06Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000010145
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Verification status:
	//
	// - **1**: Verification passed.
	//
	// - **2**: Verification failed.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetCertName() *string {
	return s.CertName
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetCertNo() *string {
	return s.CertNo
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetBizCode(v string) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.BizCode = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetCertName(v string) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.CertName = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetCertNo(v string) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.CertNo = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetCertifyId(v string) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.CertifyId = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetCreateTime(v string) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.CreateTime = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetSceneId(v int64) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.SceneId = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) SetStatus(v int32) *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos {
	s.Status = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponseBodyFaceVerifyInfos) Validate() error {
	return dara.Validate(s)
}
