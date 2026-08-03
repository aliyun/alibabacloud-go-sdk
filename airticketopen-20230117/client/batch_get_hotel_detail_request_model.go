// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetHotelDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *BatchGetHotelDetailRequest
	GetAccountNo() *int64
	SetLanguage(v string) *BatchGetHotelDetailRequest
	GetLanguage() *string
	SetStandardHotelIds(v []*string) *BatchGetHotelDetailRequest
	GetStandardHotelIds() []*string
	SetTracerId(v string) *BatchGetHotelDetailRequest
	GetTracerId() *string
}

type BatchGetHotelDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["H001","H002"]
	StandardHotelIds []*string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty" type:"Repeated"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s BatchGetHotelDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailRequest) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *BatchGetHotelDetailRequest) GetLanguage() *string {
	return s.Language
}

func (s *BatchGetHotelDetailRequest) GetStandardHotelIds() []*string {
	return s.StandardHotelIds
}

func (s *BatchGetHotelDetailRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *BatchGetHotelDetailRequest) SetAccountNo(v int64) *BatchGetHotelDetailRequest {
	s.AccountNo = &v
	return s
}

func (s *BatchGetHotelDetailRequest) SetLanguage(v string) *BatchGetHotelDetailRequest {
	s.Language = &v
	return s
}

func (s *BatchGetHotelDetailRequest) SetStandardHotelIds(v []*string) *BatchGetHotelDetailRequest {
	s.StandardHotelIds = v
	return s
}

func (s *BatchGetHotelDetailRequest) SetTracerId(v string) *BatchGetHotelDetailRequest {
	s.TracerId = &v
	return s
}

func (s *BatchGetHotelDetailRequest) Validate() error {
	return dara.Validate(s)
}
