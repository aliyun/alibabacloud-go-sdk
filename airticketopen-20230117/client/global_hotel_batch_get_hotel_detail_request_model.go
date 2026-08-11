// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelBatchGetHotelDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelBatchGetHotelDetailRequest
	GetAccountNo() *int64
	SetLanguage(v string) *GlobalHotelBatchGetHotelDetailRequest
	GetLanguage() *string
	SetStandardHotelIds(v []*string) *GlobalHotelBatchGetHotelDetailRequest
	GetStandardHotelIds() []*string
	SetTracerId(v string) *GlobalHotelBatchGetHotelDetailRequest
	GetTracerId() *string
}

type GlobalHotelBatchGetHotelDetailRequest struct {
	// The distributor account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The language (such as en or zh).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The list of standard hotel IDs. A maximum of 100 IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["H001","H002"]
	StandardHotelIds []*string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty" type:"Repeated"`
	// string
	//
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelBatchGetHotelDetailRequest) GetLanguage() *string {
	return s.Language
}

func (s *GlobalHotelBatchGetHotelDetailRequest) GetStandardHotelIds() []*string {
	return s.StandardHotelIds
}

func (s *GlobalHotelBatchGetHotelDetailRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelBatchGetHotelDetailRequest) SetAccountNo(v int64) *GlobalHotelBatchGetHotelDetailRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailRequest) SetLanguage(v string) *GlobalHotelBatchGetHotelDetailRequest {
	s.Language = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailRequest) SetStandardHotelIds(v []*string) *GlobalHotelBatchGetHotelDetailRequest {
	s.StandardHotelIds = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailRequest) SetTracerId(v string) *GlobalHotelBatchGetHotelDetailRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailRequest) Validate() error {
	return dara.Validate(s)
}
