// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubCrowdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSubCrowdsResponseBody
	GetRequestId() *string
	SetSubCrowds(v []*ListSubCrowdsResponseBodySubCrowds) *ListSubCrowdsResponseBody
	GetSubCrowds() []*ListSubCrowdsResponseBodySubCrowds
	SetTotalCount(v int64) *ListSubCrowdsResponseBody
	GetTotalCount() *int64
}

type ListSubCrowdsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C5AEB79E-FAA4-5DCE-8CD7-1CAF549ECC3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of subcrowds.
	SubCrowds []*ListSubCrowdsResponseBodySubCrowds `json:"SubCrowds,omitempty" xml:"SubCrowds,omitempty" type:"Repeated"`
	// The total number of subcrowds.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubCrowdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubCrowdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubCrowdsResponseBody) GetSubCrowds() []*ListSubCrowdsResponseBodySubCrowds {
	return s.SubCrowds
}

func (s *ListSubCrowdsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSubCrowdsResponseBody) SetRequestId(v string) *ListSubCrowdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubCrowdsResponseBody) SetSubCrowds(v []*ListSubCrowdsResponseBodySubCrowds) *ListSubCrowdsResponseBody {
	s.SubCrowds = v
	return s
}

func (s *ListSubCrowdsResponseBody) SetTotalCount(v int64) *ListSubCrowdsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubCrowdsResponseBody) Validate() error {
	if s.SubCrowds != nil {
		for _, item := range s.SubCrowds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubCrowdsResponseBodySubCrowds struct {
	// The time when the subcrowd was created.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The number of users in the subcrowd.
	//
	// example:
	//
	// 2
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The source of the subcrowd.<br>● ManualInput: Users were added manually.<br>● UploadFile: Users were added by uploading a file.<br><br>
	//
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The subcrowd ID.
	//
	// example:
	//
	// 3
	SubCrowdId *string `json:"SubCrowdId,omitempty" xml:"SubCrowdId,omitempty"`
	// The users in the subcrowd.
	//
	// example:
	//
	// user1,user2
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ListSubCrowdsResponseBodySubCrowds) String() string {
	return dara.Prettify(s)
}

func (s ListSubCrowdsResponseBodySubCrowds) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsResponseBodySubCrowds) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListSubCrowdsResponseBodySubCrowds) GetQuantity() *int32 {
	return s.Quantity
}

func (s *ListSubCrowdsResponseBodySubCrowds) GetSource() *string {
	return s.Source
}

func (s *ListSubCrowdsResponseBodySubCrowds) GetSubCrowdId() *string {
	return s.SubCrowdId
}

func (s *ListSubCrowdsResponseBodySubCrowds) GetUsers() *string {
	return s.Users
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetGmtCreateTime(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetQuantity(v int32) *ListSubCrowdsResponseBodySubCrowds {
	s.Quantity = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetSource(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.Source = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetSubCrowdId(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.SubCrowdId = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) SetUsers(v string) *ListSubCrowdsResponseBodySubCrowds {
	s.Users = &v
	return s
}

func (s *ListSubCrowdsResponseBodySubCrowds) Validate() error {
	return dara.Validate(s)
}
