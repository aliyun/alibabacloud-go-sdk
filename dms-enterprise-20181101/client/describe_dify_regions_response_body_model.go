// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDifyRegionsResponseBody
	GetCode() *string
	SetData(v []*DescribeDifyRegionsResponseBodyData) *DescribeDifyRegionsResponseBody
	GetData() []*DescribeDifyRegionsResponseBodyData
	SetErrorCode(v string) *DescribeDifyRegionsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DescribeDifyRegionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDifyRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDifyRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDifyRegionsResponseBody
	GetSuccess() *bool
}

type DescribeDifyRegionsResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeDifyRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode      *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDifyRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDifyRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDifyRegionsResponseBody) GetData() []*DescribeDifyRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeDifyRegionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDifyRegionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDifyRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDifyRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDifyRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDifyRegionsResponseBody) SetCode(v string) *DescribeDifyRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDifyRegionsResponseBody) SetData(v []*DescribeDifyRegionsResponseBodyData) *DescribeDifyRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDifyRegionsResponseBody) SetErrorCode(v string) *DescribeDifyRegionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDifyRegionsResponseBody) SetHttpStatusCode(v int32) *DescribeDifyRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDifyRegionsResponseBody) SetMessage(v string) *DescribeDifyRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDifyRegionsResponseBody) SetRequestId(v string) *DescribeDifyRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDifyRegionsResponseBody) SetSuccess(v bool) *DescribeDifyRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDifyRegionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDifyRegionsResponseBodyData struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDifyRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDifyRegionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDifyRegionsResponseBodyData) SetRegionId(v string) *DescribeDifyRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDifyRegionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
