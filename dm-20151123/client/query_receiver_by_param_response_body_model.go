// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReceiverByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextStart(v string) *QueryReceiverByParamResponseBody
	GetNextStart() *string
	SetPageSize(v int32) *QueryReceiverByParamResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryReceiverByParamResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryReceiverByParamResponseBody
	GetTotalCount() *int32
	SetData(v *QueryReceiverByParamResponseBodyData) *QueryReceiverByParamResponseBody
	GetData() *QueryReceiverByParamResponseBodyData
}

type QueryReceiverByParamResponseBody struct {
	// Used for paging. If more results are available, set this value as the NextStart parameter in your next request.
	//
	// example:
	//
	// 6aec200853#102#1638894326#test@example.com
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// Number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of entries.
	//
	// example:
	//
	// 15
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryReceiverByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryReceiverByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBody) GetNextStart() *string {
	return s.NextStart
}

func (s *QueryReceiverByParamResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryReceiverByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReceiverByParamResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryReceiverByParamResponseBody) GetData() *QueryReceiverByParamResponseBodyData {
	return s.Data
}

func (s *QueryReceiverByParamResponseBody) SetNextStart(v string) *QueryReceiverByParamResponseBody {
	s.NextStart = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetPageSize(v int32) *QueryReceiverByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetRequestId(v string) *QueryReceiverByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetTotalCount(v int32) *QueryReceiverByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetData(v *QueryReceiverByParamResponseBodyData) *QueryReceiverByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryReceiverByParamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryReceiverByParamResponseBodyData struct {
	Receiver []*QueryReceiverByParamResponseBodyDataReceiver `json:"receiver,omitempty" xml:"receiver,omitempty" type:"Repeated"`
}

func (s QueryReceiverByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBodyData) GetReceiver() []*QueryReceiverByParamResponseBodyDataReceiver {
	return s.Receiver
}

func (s *QueryReceiverByParamResponseBodyData) SetReceiver(v []*QueryReceiverByParamResponseBodyDataReceiver) *QueryReceiverByParamResponseBodyData {
	s.Receiver = v
	return s
}

func (s *QueryReceiverByParamResponseBodyData) Validate() error {
	if s.Receiver != nil {
		for _, item := range s.Receiver {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryReceiverByParamResponseBodyDataReceiver struct {
	Count           *string `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Desc            *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	ReceiverId      *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ReceiversAlias  *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	ReceiversName   *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ReceiversStatus *string `json:"ReceiversStatus,omitempty" xml:"ReceiversStatus,omitempty"`
	UtcCreateTime   *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryReceiverByParamResponseBodyDataReceiver) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverByParamResponseBodyDataReceiver) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetCount() *string {
	return s.Count
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetDesc() *string {
	return s.Desc
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetReceiversAlias() *string {
	return s.ReceiversAlias
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetReceiversName() *string {
	return s.ReceiversName
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetReceiversStatus() *string {
	return s.ReceiversStatus
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetCount(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.Count = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetCreateTime(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetDesc(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.Desc = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiverId(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiverId = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversAlias(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversAlias = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversName(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversName = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversStatus(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversStatus = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetUtcCreateTime(v int64) *QueryReceiverByParamResponseBodyDataReceiver {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) Validate() error {
	return dara.Validate(s)
}
