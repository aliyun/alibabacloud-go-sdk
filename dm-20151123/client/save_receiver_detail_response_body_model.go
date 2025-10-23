// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveReceiverDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SaveReceiverDetailResponseBodyData) *SaveReceiverDetailResponseBody
	GetData() *SaveReceiverDetailResponseBodyData
	SetErrorCount(v int32) *SaveReceiverDetailResponseBody
	GetErrorCount() *int32
	SetRequestId(v string) *SaveReceiverDetailResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *SaveReceiverDetailResponseBody
	GetSuccessCount() *int32
}

type SaveReceiverDetailResponseBody struct {
	// List of recipient addresses that failed to upload.
	Data *SaveReceiverDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Number of errors.
	//
	// example:
	//
	// 638
	ErrorCount *int32 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Number of successes.
	//
	// example:
	//
	// 274
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s SaveReceiverDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBody) GetData() *SaveReceiverDetailResponseBodyData {
	return s.Data
}

func (s *SaveReceiverDetailResponseBody) GetErrorCount() *int32 {
	return s.ErrorCount
}

func (s *SaveReceiverDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveReceiverDetailResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *SaveReceiverDetailResponseBody) SetData(v *SaveReceiverDetailResponseBodyData) *SaveReceiverDetailResponseBody {
	s.Data = v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetErrorCount(v int32) *SaveReceiverDetailResponseBody {
	s.ErrorCount = &v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetRequestId(v string) *SaveReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetSuccessCount(v int32) *SaveReceiverDetailResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *SaveReceiverDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveReceiverDetailResponseBodyData struct {
	Detail []*SaveReceiverDetailResponseBodyDataDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s SaveReceiverDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SaveReceiverDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBodyData) GetDetail() []*SaveReceiverDetailResponseBodyDataDetail {
	return s.Detail
}

func (s *SaveReceiverDetailResponseBodyData) SetDetail(v []*SaveReceiverDetailResponseBodyDataDetail) *SaveReceiverDetailResponseBodyData {
	s.Detail = v
	return s
}

func (s *SaveReceiverDetailResponseBodyData) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveReceiverDetailResponseBodyDataDetail struct {
	// Recipient address.
	//
	// example:
	//
	// test@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// XXX
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s SaveReceiverDetailResponseBodyDataDetail) String() string {
	return dara.Prettify(s)
}

func (s SaveReceiverDetailResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBodyDataDetail) GetEmail() *string {
	return s.Email
}

func (s *SaveReceiverDetailResponseBodyDataDetail) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SaveReceiverDetailResponseBodyDataDetail) SetEmail(v string) *SaveReceiverDetailResponseBodyDataDetail {
	s.Email = &v
	return s
}

func (s *SaveReceiverDetailResponseBodyDataDetail) SetErrMessage(v string) *SaveReceiverDetailResponseBodyDataDetail {
	s.ErrMessage = &v
	return s
}

func (s *SaveReceiverDetailResponseBodyDataDetail) Validate() error {
	return dara.Validate(s)
}
