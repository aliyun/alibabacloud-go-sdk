// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyEditionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDifyEditionsResponseBody
	GetCode() *string
	SetData(v *DescribeDifyEditionsResponseBodyData) *DescribeDifyEditionsResponseBody
	GetData() *DescribeDifyEditionsResponseBodyData
	SetErrorCode(v string) *DescribeDifyEditionsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DescribeDifyEditionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDifyEditionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDifyEditionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDifyEditionsResponseBody
	GetSuccess() *bool
}

type DescribeDifyEditionsResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeDifyEditionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDifyEditionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyEditionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDifyEditionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDifyEditionsResponseBody) GetData() *DescribeDifyEditionsResponseBodyData {
	return s.Data
}

func (s *DescribeDifyEditionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDifyEditionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDifyEditionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDifyEditionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDifyEditionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDifyEditionsResponseBody) SetCode(v string) *DescribeDifyEditionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDifyEditionsResponseBody) SetData(v *DescribeDifyEditionsResponseBodyData) *DescribeDifyEditionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDifyEditionsResponseBody) SetErrorCode(v string) *DescribeDifyEditionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDifyEditionsResponseBody) SetHttpStatusCode(v int32) *DescribeDifyEditionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDifyEditionsResponseBody) SetMessage(v string) *DescribeDifyEditionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDifyEditionsResponseBody) SetRequestId(v string) *DescribeDifyEditionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDifyEditionsResponseBody) SetSuccess(v bool) *DescribeDifyEditionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDifyEditionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDifyEditionsResponseBodyData struct {
	Community     []*string `json:"Community,omitempty" xml:"Community,omitempty" type:"Repeated"`
	Enterprise    []*string `json:"Enterprise,omitempty" xml:"Enterprise,omitempty" type:"Repeated"`
	OpenCommunity []*string `json:"OpenCommunity,omitempty" xml:"OpenCommunity,omitempty" type:"Repeated"`
}

func (s DescribeDifyEditionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyEditionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDifyEditionsResponseBodyData) GetCommunity() []*string {
	return s.Community
}

func (s *DescribeDifyEditionsResponseBodyData) GetEnterprise() []*string {
	return s.Enterprise
}

func (s *DescribeDifyEditionsResponseBodyData) GetOpenCommunity() []*string {
	return s.OpenCommunity
}

func (s *DescribeDifyEditionsResponseBodyData) SetCommunity(v []*string) *DescribeDifyEditionsResponseBodyData {
	s.Community = v
	return s
}

func (s *DescribeDifyEditionsResponseBodyData) SetEnterprise(v []*string) *DescribeDifyEditionsResponseBodyData {
	s.Enterprise = v
	return s
}

func (s *DescribeDifyEditionsResponseBodyData) SetOpenCommunity(v []*string) *DescribeDifyEditionsResponseBodyData {
	s.OpenCommunity = v
	return s
}

func (s *DescribeDifyEditionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
