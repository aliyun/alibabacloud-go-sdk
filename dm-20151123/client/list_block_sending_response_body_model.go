// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBlockSendingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListBlockSendingResponseBodyData) *ListBlockSendingResponseBody
	GetData() []*ListBlockSendingResponseBodyData
	SetMaxResults(v int32) *ListBlockSendingResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListBlockSendingResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBlockSendingResponseBody
	GetRequestId() *string
}

type ListBlockSendingResponseBody struct {
	Data []*ListBlockSendingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// xxxxyyyy
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBlockSendingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBlockSendingResponseBody) GoString() string {
	return s.String()
}

func (s *ListBlockSendingResponseBody) GetData() []*ListBlockSendingResponseBodyData {
	return s.Data
}

func (s *ListBlockSendingResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBlockSendingResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBlockSendingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBlockSendingResponseBody) SetData(v []*ListBlockSendingResponseBodyData) *ListBlockSendingResponseBody {
	s.Data = v
	return s
}

func (s *ListBlockSendingResponseBody) SetMaxResults(v int32) *ListBlockSendingResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBlockSendingResponseBody) SetNextToken(v string) *ListBlockSendingResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBlockSendingResponseBody) SetRequestId(v string) *ListBlockSendingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBlockSendingResponseBody) Validate() error {
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

type ListBlockSendingResponseBodyData struct {
	// example:
	//
	// xxxx@rcpt.com
	BlockEmail *string `json:"BlockEmail,omitempty" xml:"BlockEmail,omitempty"`
	// example:
	//
	// 1723259364
	BlockTime *int32 `json:"BlockTime,omitempty" xml:"BlockTime,omitempty"`
	// example:
	//
	// 1
	Reason *int32 `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 1723249364
	SendTime *int32 `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// example:
	//
	// xxxx@sender.com
	SenderEmail *string `json:"SenderEmail,omitempty" xml:"SenderEmail,omitempty"`
}

func (s ListBlockSendingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBlockSendingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBlockSendingResponseBodyData) GetBlockEmail() *string {
	return s.BlockEmail
}

func (s *ListBlockSendingResponseBodyData) GetBlockTime() *int32 {
	return s.BlockTime
}

func (s *ListBlockSendingResponseBodyData) GetReason() *int32 {
	return s.Reason
}

func (s *ListBlockSendingResponseBodyData) GetSendTime() *int32 {
	return s.SendTime
}

func (s *ListBlockSendingResponseBodyData) GetSenderEmail() *string {
	return s.SenderEmail
}

func (s *ListBlockSendingResponseBodyData) SetBlockEmail(v string) *ListBlockSendingResponseBodyData {
	s.BlockEmail = &v
	return s
}

func (s *ListBlockSendingResponseBodyData) SetBlockTime(v int32) *ListBlockSendingResponseBodyData {
	s.BlockTime = &v
	return s
}

func (s *ListBlockSendingResponseBodyData) SetReason(v int32) *ListBlockSendingResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListBlockSendingResponseBodyData) SetSendTime(v int32) *ListBlockSendingResponseBodyData {
	s.SendTime = &v
	return s
}

func (s *ListBlockSendingResponseBodyData) SetSenderEmail(v string) *ListBlockSendingResponseBodyData {
	s.SenderEmail = &v
	return s
}

func (s *ListBlockSendingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
