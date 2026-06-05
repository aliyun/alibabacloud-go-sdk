// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDatasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyUserDatasResponseBody
	GetCode() *int64
	SetMessage(v string) *DescribeComfyUserDatasResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeComfyUserDatasResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyUserDatasResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeComfyUserDatasResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeComfyUserDatasResponseBody
	GetTotal() *int32
	SetUserDatas(v []*DescribeComfyUserDatasResponseBodyUserDatas) *DescribeComfyUserDatasResponseBody
	GetUserDatas() []*DescribeComfyUserDatasResponseBodyUserDatas
}

type DescribeComfyUserDatasResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Total     *int32                                         `json:"Total,omitempty" xml:"Total,omitempty"`
	UserDatas []*DescribeComfyUserDatasResponseBodyUserDatas `json:"UserDatas,omitempty" xml:"UserDatas,omitempty" type:"Repeated"`
}

func (s DescribeComfyUserDatasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDatasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDatasResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyUserDatasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyUserDatasResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyUserDatasResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyUserDatasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyUserDatasResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeComfyUserDatasResponseBody) GetUserDatas() []*DescribeComfyUserDatasResponseBodyUserDatas {
	return s.UserDatas
}

func (s *DescribeComfyUserDatasResponseBody) SetCode(v int64) *DescribeComfyUserDatasResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) SetMessage(v string) *DescribeComfyUserDatasResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) SetPageNumber(v int32) *DescribeComfyUserDatasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) SetPageSize(v int32) *DescribeComfyUserDatasResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) SetRequestId(v string) *DescribeComfyUserDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) SetTotal(v int32) *DescribeComfyUserDatasResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) SetUserDatas(v []*DescribeComfyUserDatasResponseBodyUserDatas) *DescribeComfyUserDatasResponseBody {
	s.UserDatas = v
	return s
}

func (s *DescribeComfyUserDatasResponseBody) Validate() error {
	if s.UserDatas != nil {
		for _, item := range s.UserDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComfyUserDatasResponseBodyUserDatas struct {
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1024
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1776646928000
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeComfyUserDatasResponseBodyUserDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDatasResponseBodyUserDatas) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) GetFileName() *string {
	return s.FileName
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) GetType() *string {
	return s.Type
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) SetFileName(v string) *DescribeComfyUserDatasResponseBodyUserDatas {
	s.FileName = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) SetFileSize(v int64) *DescribeComfyUserDatasResponseBodyUserDatas {
	s.FileSize = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) SetType(v string) *DescribeComfyUserDatasResponseBodyUserDatas {
	s.Type = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) SetUpdatedTime(v string) *DescribeComfyUserDatasResponseBodyUserDatas {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeComfyUserDatasResponseBodyUserDatas) Validate() error {
	return dara.Validate(s)
}
