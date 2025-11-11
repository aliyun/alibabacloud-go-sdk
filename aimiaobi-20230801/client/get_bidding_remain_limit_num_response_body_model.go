// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBiddingRemainLimitNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBiddingRemainLimitNumResponseBody
	GetCode() *string
	SetData(v *GetBiddingRemainLimitNumResponseBodyData) *GetBiddingRemainLimitNumResponseBody
	GetData() *GetBiddingRemainLimitNumResponseBodyData
	SetHttpStatusCode(v int32) *GetBiddingRemainLimitNumResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBiddingRemainLimitNumResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBiddingRemainLimitNumResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBiddingRemainLimitNumResponseBody
	GetSuccess() *bool
}

type GetBiddingRemainLimitNumResponseBody struct {
	// example:
	//
	// successful
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBiddingRemainLimitNumResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBiddingRemainLimitNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingRemainLimitNumResponseBody) GoString() string {
	return s.String()
}

func (s *GetBiddingRemainLimitNumResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBiddingRemainLimitNumResponseBody) GetData() *GetBiddingRemainLimitNumResponseBodyData {
	return s.Data
}

func (s *GetBiddingRemainLimitNumResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBiddingRemainLimitNumResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBiddingRemainLimitNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBiddingRemainLimitNumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBiddingRemainLimitNumResponseBody) SetCode(v string) *GetBiddingRemainLimitNumResponseBody {
	s.Code = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBody) SetData(v *GetBiddingRemainLimitNumResponseBodyData) *GetBiddingRemainLimitNumResponseBody {
	s.Data = v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBody) SetHttpStatusCode(v int32) *GetBiddingRemainLimitNumResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBody) SetMessage(v string) *GetBiddingRemainLimitNumResponseBody {
	s.Message = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBody) SetRequestId(v string) *GetBiddingRemainLimitNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBody) SetSuccess(v bool) *GetBiddingRemainLimitNumResponseBody {
	s.Success = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBiddingRemainLimitNumResponseBodyData struct {
	// example:
	//
	// 1
	RemainNum *int32 `json:"RemainNum,omitempty" xml:"RemainNum,omitempty"`
}

func (s GetBiddingRemainLimitNumResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingRemainLimitNumResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBiddingRemainLimitNumResponseBodyData) GetRemainNum() *int32 {
	return s.RemainNum
}

func (s *GetBiddingRemainLimitNumResponseBodyData) SetRemainNum(v int32) *GetBiddingRemainLimitNumResponseBodyData {
	s.RemainNum = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponseBodyData) Validate() error {
	return dara.Validate(s)
}
