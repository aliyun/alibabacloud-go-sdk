// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetHotelDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *BatchGetHotelDetailShrinkRequest
	GetAccountNo() *int64
	SetLanguage(v string) *BatchGetHotelDetailShrinkRequest
	GetLanguage() *string
	SetStandardHotelIdsShrink(v string) *BatchGetHotelDetailShrinkRequest
	GetStandardHotelIdsShrink() *string
	SetTracerId(v string) *BatchGetHotelDetailShrinkRequest
	GetTracerId() *string
}

type BatchGetHotelDetailShrinkRequest struct {
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
	StandardHotelIdsShrink *string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s BatchGetHotelDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *BatchGetHotelDetailShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *BatchGetHotelDetailShrinkRequest) GetStandardHotelIdsShrink() *string {
	return s.StandardHotelIdsShrink
}

func (s *BatchGetHotelDetailShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *BatchGetHotelDetailShrinkRequest) SetAccountNo(v int64) *BatchGetHotelDetailShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *BatchGetHotelDetailShrinkRequest) SetLanguage(v string) *BatchGetHotelDetailShrinkRequest {
	s.Language = &v
	return s
}

func (s *BatchGetHotelDetailShrinkRequest) SetStandardHotelIdsShrink(v string) *BatchGetHotelDetailShrinkRequest {
	s.StandardHotelIdsShrink = &v
	return s
}

func (s *BatchGetHotelDetailShrinkRequest) SetTracerId(v string) *BatchGetHotelDetailShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *BatchGetHotelDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
