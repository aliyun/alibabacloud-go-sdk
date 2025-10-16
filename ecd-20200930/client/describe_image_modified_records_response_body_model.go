// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageModifiedRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageModifiedRecords(v []*DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) *DescribeImageModifiedRecordsResponseBody
	GetImageModifiedRecords() []*DescribeImageModifiedRecordsResponseBodyImageModifiedRecords
	SetNextToken(v string) *DescribeImageModifiedRecordsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeImageModifiedRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImageModifiedRecordsResponseBody
	GetTotalCount() *int32
}

type DescribeImageModifiedRecordsResponseBody struct {
	// The image change records.
	ImageModifiedRecords []*DescribeImageModifiedRecordsResponseBodyImageModifiedRecords `json:"ImageModifiedRecords,omitempty" xml:"ImageModifiedRecords,omitempty" type:"Repeated"`
	// If the NextToken parameter is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6koN7RqHg3d2z8LKmSoe821
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC40EE61-7E83-59ED-AEA6-7EE9C437F352
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of image modification records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageModifiedRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModifiedRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageModifiedRecordsResponseBody) GetImageModifiedRecords() []*DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	return s.ImageModifiedRecords
}

func (s *DescribeImageModifiedRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImageModifiedRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageModifiedRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageModifiedRecordsResponseBody) SetImageModifiedRecords(v []*DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) *DescribeImageModifiedRecordsResponseBody {
	s.ImageModifiedRecords = v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBody) SetNextToken(v string) *DescribeImageModifiedRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBody) SetRequestId(v string) *DescribeImageModifiedRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBody) SetTotalCount(v int32) *DescribeImageModifiedRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBody) Validate() error {
	if s.ImageModifiedRecords != nil {
		for _, item := range s.ImageModifiedRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageModifiedRecordsResponseBodyImageModifiedRecords struct {
	// The ID of the original image.
	//
	// example:
	//
	// m-8rnz2imrpcfuh****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the original image.
	//
	// example:
	//
	// win10-0307
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the new image after the image was modified.
	//
	// example:
	//
	// xxxNewImageID
	NewImageId *string `json:"NewImageId,omitempty" xml:"NewImageId,omitempty"`
	// The name of the new image after the image was modified.
	//
	// example:
	//
	// xxxxImageID
	NewImageName *string `json:"NewImageName,omitempty" xml:"NewImageName,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the image modification.
	//
	// Valid values:
	//
	// 	- 0: The image is being modified.
	//
	// 	- 1: The image is successfully modified.
	//
	// 	- 2: The image fails to be modified.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the image was last modified.
	//
	// example:
	//
	// 2022-03-03T02:43:44.851Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GoString() string {
	return s.String()
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetNewImageId() *string {
	return s.NewImageId
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetNewImageName() *string {
	return s.NewImageName
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetReason() *string {
	return s.Reason
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetImageId(v string) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.ImageId = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetImageName(v string) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.ImageName = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetNewImageId(v string) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.NewImageId = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetNewImageName(v string) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.NewImageName = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetReason(v string) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.Reason = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetStatus(v int32) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.Status = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) SetUpdateTime(v string) *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords {
	s.UpdateTime = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponseBodyImageModifiedRecords) Validate() error {
	return dara.Validate(s)
}
