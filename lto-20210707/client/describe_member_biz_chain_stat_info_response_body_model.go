// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberBizChainStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMemberBizChainStatInfoResponseBody
	GetCode() *string
	SetData(v []*DescribeMemberBizChainStatInfoResponseBodyData) *DescribeMemberBizChainStatInfoResponseBody
	GetData() []*DescribeMemberBizChainStatInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeMemberBizChainStatInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeMemberBizChainStatInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMemberBizChainStatInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMemberBizChainStatInfoResponseBody
	GetSuccess() *bool
}

type DescribeMemberBizChainStatInfoResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeMemberBizChainStatInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMemberBizChainStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberBizChainStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberBizChainStatInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMemberBizChainStatInfoResponseBody) GetData() []*DescribeMemberBizChainStatInfoResponseBodyData {
	return s.Data
}

func (s *DescribeMemberBizChainStatInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeMemberBizChainStatInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMemberBizChainStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMemberBizChainStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMemberBizChainStatInfoResponseBody) SetCode(v string) *DescribeMemberBizChainStatInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBody) SetData(v []*DescribeMemberBizChainStatInfoResponseBodyData) *DescribeMemberBizChainStatInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBody) SetHttpStatusCode(v int32) *DescribeMemberBizChainStatInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBody) SetMessage(v string) *DescribeMemberBizChainStatInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBody) SetRequestId(v string) *DescribeMemberBizChainStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBody) SetSuccess(v bool) *DescribeMemberBizChainStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBody) Validate() error {
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

type DescribeMemberBizChainStatInfoResponseBodyData struct {
	BizChainName *string `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	UsedCount    *int64  `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeMemberBizChainStatInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberBizChainStatInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMemberBizChainStatInfoResponseBodyData) GetBizChainName() *string {
	return s.BizChainName
}

func (s *DescribeMemberBizChainStatInfoResponseBodyData) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeMemberBizChainStatInfoResponseBodyData) SetBizChainName(v string) *DescribeMemberBizChainStatInfoResponseBodyData {
	s.BizChainName = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBodyData) SetUsedCount(v int64) *DescribeMemberBizChainStatInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *DescribeMemberBizChainStatInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
