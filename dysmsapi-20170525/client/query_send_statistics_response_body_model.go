// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySendStatisticsResponseBody
	GetCode() *string
	SetData(v *QuerySendStatisticsResponseBodyData) *QuerySendStatisticsResponseBody
	GetData() *QuerySendStatisticsResponseBodyData
	SetMessage(v string) *QuerySendStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySendStatisticsResponseBody
	GetRequestId() *string
}

type QuerySendStatisticsResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QuerySendStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E47708****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySendStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySendStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySendStatisticsResponseBody) GetData() *QuerySendStatisticsResponseBodyData {
	return s.Data
}

func (s *QuerySendStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySendStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySendStatisticsResponseBody) SetCode(v string) *QuerySendStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySendStatisticsResponseBody) SetData(v *QuerySendStatisticsResponseBodyData) *QuerySendStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *QuerySendStatisticsResponseBody) SetMessage(v string) *QuerySendStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySendStatisticsResponseBody) SetRequestId(v string) *QuerySendStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySendStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySendStatisticsResponseBodyData struct {
	// The details of the data returned.
	TargetList []*QuerySendStatisticsResponseBodyDataTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s QuerySendStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySendStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponseBodyData) GetTargetList() []*QuerySendStatisticsResponseBodyDataTargetList {
	return s.TargetList
}

func (s *QuerySendStatisticsResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *QuerySendStatisticsResponseBodyData) SetTargetList(v []*QuerySendStatisticsResponseBodyDataTargetList) *QuerySendStatisticsResponseBodyData {
	s.TargetList = v
	return s
}

func (s *QuerySendStatisticsResponseBodyData) SetTotalSize(v int64) *QuerySendStatisticsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyData) Validate() error {
	if s.TargetList != nil {
		for _, item := range s.TargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySendStatisticsResponseBodyDataTargetList struct {
	// The number of messages without a delivery receipt.
	//
	// example:
	//
	// 1
	NoRespondedCount *int64 `json:"NoRespondedCount,omitempty" xml:"NoRespondedCount,omitempty"`
	// The number of messages with a delivery receipt that indicates a failure.
	//
	// example:
	//
	// 2
	RespondedFailCount *int64 `json:"RespondedFailCount,omitempty" xml:"RespondedFailCount,omitempty"`
	// The number of messages with a delivery receipt that indicates a success.
	//
	// example:
	//
	// 17
	RespondedSuccessCount *int64 `json:"RespondedSuccessCount,omitempty" xml:"RespondedSuccessCount,omitempty"`
	// The date when the message is sent. Format: yyyyMMdd. Example: 20181225.
	//
	// example:
	//
	// 20201010
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The number of delivered messages.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySendStatisticsResponseBodyDataTargetList) String() string {
	return dara.Prettify(s)
}

func (s QuerySendStatisticsResponseBodyDataTargetList) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) GetNoRespondedCount() *int64 {
	return s.NoRespondedCount
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) GetRespondedFailCount() *int64 {
	return s.RespondedFailCount
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) GetRespondedSuccessCount() *int64 {
	return s.RespondedSuccessCount
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) GetSendDate() *string {
	return s.SendDate
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetNoRespondedCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.NoRespondedCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetRespondedFailCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.RespondedFailCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetRespondedSuccessCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.RespondedSuccessCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetSendDate(v string) *QuerySendStatisticsResponseBodyDataTargetList {
	s.SendDate = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetTotalCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.TotalCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) Validate() error {
	return dara.Validate(s)
}
