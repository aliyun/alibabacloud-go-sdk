// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelBatchGetHotelDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelBatchGetHotelDetailShrinkRequest
	GetAccountNo() *int64
	SetLanguage(v string) *GlobalHotelBatchGetHotelDetailShrinkRequest
	GetLanguage() *string
	SetStandardHotelIdsShrink(v string) *GlobalHotelBatchGetHotelDetailShrinkRequest
	GetStandardHotelIdsShrink() *string
	SetTracerId(v string) *GlobalHotelBatchGetHotelDetailShrinkRequest
	GetTracerId() *string
}

type GlobalHotelBatchGetHotelDetailShrinkRequest struct {
	// The ID of the distributor account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The language. For example, en or zh.
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
	StandardHotelIdsShrink *string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty"`
	// string
	//
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) GetStandardHotelIdsShrink() *string {
	return s.StandardHotelIdsShrink
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) SetAccountNo(v int64) *GlobalHotelBatchGetHotelDetailShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) SetLanguage(v string) *GlobalHotelBatchGetHotelDetailShrinkRequest {
	s.Language = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) SetStandardHotelIdsShrink(v string) *GlobalHotelBatchGetHotelDetailShrinkRequest {
	s.StandardHotelIdsShrink = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) SetTracerId(v string) *GlobalHotelBatchGetHotelDetailShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
