// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBizChainStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeBizChainStatInfoResponseBody
	GetCode() *string
	SetData(v []*DescribeBizChainStatInfoResponseBodyData) *DescribeBizChainStatInfoResponseBody
	GetData() []*DescribeBizChainStatInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeBizChainStatInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeBizChainStatInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBizChainStatInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBizChainStatInfoResponseBody
	GetSuccess() *bool
}

type DescribeBizChainStatInfoResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeBizChainStatInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBizChainStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizChainStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizChainStatInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBizChainStatInfoResponseBody) GetData() []*DescribeBizChainStatInfoResponseBodyData {
	return s.Data
}

func (s *DescribeBizChainStatInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBizChainStatInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBizChainStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBizChainStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBizChainStatInfoResponseBody) SetCode(v string) *DescribeBizChainStatInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBody) SetData(v []*DescribeBizChainStatInfoResponseBodyData) *DescribeBizChainStatInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBizChainStatInfoResponseBody) SetHttpStatusCode(v int32) *DescribeBizChainStatInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBody) SetMessage(v string) *DescribeBizChainStatInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBody) SetRequestId(v string) *DescribeBizChainStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBody) SetSuccess(v bool) *DescribeBizChainStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBody) Validate() error {
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

type DescribeBizChainStatInfoResponseBodyData struct {
	BizChainName *string `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	UsedCount    *int64  `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeBizChainStatInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizChainStatInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBizChainStatInfoResponseBodyData) GetBizChainName() *string {
	return s.BizChainName
}

func (s *DescribeBizChainStatInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeBizChainStatInfoResponseBodyData) SetBizChainName(v string) *DescribeBizChainStatInfoResponseBodyData {
	s.BizChainName = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBodyData) SetUsedCount(v int64) *DescribeBizChainStatInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeBizChainStatInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
