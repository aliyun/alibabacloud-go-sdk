// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyProductionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyProductionsResponseBody
	GetCode() *int64
	SetMessage(v string) *DescribeComfyProductionsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeComfyProductionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyProductionsResponseBody
	GetPageSize() *int32
	SetProductions(v []*DescribeComfyProductionsResponseBodyProductions) *DescribeComfyProductionsResponseBody
	GetProductions() []*DescribeComfyProductionsResponseBodyProductions
	SetRequestId(v string) *DescribeComfyProductionsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeComfyProductionsResponseBody
	GetTotal() *int32
}

type DescribeComfyProductionsResponseBody struct {
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
	PageSize    *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Productions []*DescribeComfyProductionsResponseBodyProductions `json:"Productions,omitempty" xml:"Productions,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeComfyProductionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyProductionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyProductionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyProductionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyProductionsResponseBody) GetProductions() []*DescribeComfyProductionsResponseBodyProductions {
	return s.Productions
}

func (s *DescribeComfyProductionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyProductionsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeComfyProductionsResponseBody) SetCode(v int64) *DescribeComfyProductionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyProductionsResponseBody) SetMessage(v string) *DescribeComfyProductionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyProductionsResponseBody) SetPageNumber(v int32) *DescribeComfyProductionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyProductionsResponseBody) SetPageSize(v int32) *DescribeComfyProductionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyProductionsResponseBody) SetProductions(v []*DescribeComfyProductionsResponseBodyProductions) *DescribeComfyProductionsResponseBody {
	s.Productions = v
	return s
}

func (s *DescribeComfyProductionsResponseBody) SetRequestId(v string) *DescribeComfyProductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyProductionsResponseBody) SetTotal(v int32) *DescribeComfyProductionsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeComfyProductionsResponseBody) Validate() error {
	if s.Productions != nil {
		for _, item := range s.Productions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComfyProductionsResponseBodyProductions struct {
	// example:
	//
	// 6c8234f4-d1e1-4cea-b08b-7926fbdea144
	ComfyTaskId *string `json:"ComfyTaskId,omitempty" xml:"ComfyTaskId,omitempty"`
	// example:
	//
	// 2026-02-06T20:20:26+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 1755051607877.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 3e5bda20-5cd4-4d55-8d23-88d624a18caa
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1773707865
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeComfyProductionsResponseBodyProductions) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionsResponseBodyProductions) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionsResponseBodyProductions) GetComfyTaskId() *string {
	return s.ComfyTaskId
}

func (s *DescribeComfyProductionsResponseBodyProductions) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeComfyProductionsResponseBodyProductions) GetFileName() *string {
	return s.FileName
}

func (s *DescribeComfyProductionsResponseBodyProductions) GetProductionId() *string {
	return s.ProductionId
}

func (s *DescribeComfyProductionsResponseBodyProductions) GetState() *string {
	return s.State
}

func (s *DescribeComfyProductionsResponseBodyProductions) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *DescribeComfyProductionsResponseBodyProductions) SetComfyTaskId(v string) *DescribeComfyProductionsResponseBodyProductions {
	s.ComfyTaskId = &v
	return s
}

func (s *DescribeComfyProductionsResponseBodyProductions) SetCreationTime(v string) *DescribeComfyProductionsResponseBodyProductions {
	s.CreationTime = &v
	return s
}

func (s *DescribeComfyProductionsResponseBodyProductions) SetFileName(v string) *DescribeComfyProductionsResponseBodyProductions {
	s.FileName = &v
	return s
}

func (s *DescribeComfyProductionsResponseBodyProductions) SetProductionId(v string) *DescribeComfyProductionsResponseBodyProductions {
	s.ProductionId = &v
	return s
}

func (s *DescribeComfyProductionsResponseBodyProductions) SetState(v string) *DescribeComfyProductionsResponseBodyProductions {
	s.State = &v
	return s
}

func (s *DescribeComfyProductionsResponseBodyProductions) SetUpdatedTime(v string) *DescribeComfyProductionsResponseBodyProductions {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeComfyProductionsResponseBodyProductions) Validate() error {
	return dara.Validate(s)
}
