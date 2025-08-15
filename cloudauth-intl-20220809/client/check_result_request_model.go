// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraImageControlList(v string) *CheckResultRequest
	GetExtraImageControlList() *string
	SetIsReturnImage(v string) *CheckResultRequest
	GetIsReturnImage() *string
	SetMerchantBizId(v string) *CheckResultRequest
	GetMerchantBizId() *string
	SetReturnFiveCategorySpoofResult(v string) *CheckResultRequest
	GetReturnFiveCategorySpoofResult() *string
	SetTransactionId(v string) *CheckResultRequest
	GetTransactionId() *string
}

type CheckResultRequest struct {
	// Return additional information.
	//
	// example:
	//
	// ***
	ExtraImageControlList *string `json:"ExtraImageControlList,omitempty" xml:"ExtraImageControlList,omitempty"`
	// Whether to return images.
	//
	// - Y: Return
	//
	// - N: Do not return
	//
	// example:
	//
	// N
	IsReturnImage *string `json:"IsReturnImage,omitempty" xml:"IsReturnImage,omitempty"`
	// A unique business identifier defined by the merchant, used for subsequent troubleshooting. It supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure its uniqueness.
	//
	// example:
	//
	// djs20d***9-dsskc
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Whether to return anti-fraud detection results.
	//
	// example:
	//
	// Y
	ReturnFiveCategorySpoofResult *string `json:"ReturnFiveCategorySpoofResult,omitempty" xml:"ReturnFiveCategorySpoofResult,omitempty"`
	// Authentication ID.
	//
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckResultRequest) GoString() string {
	return s.String()
}

func (s *CheckResultRequest) GetExtraImageControlList() *string {
	return s.ExtraImageControlList
}

func (s *CheckResultRequest) GetIsReturnImage() *string {
	return s.IsReturnImage
}

func (s *CheckResultRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *CheckResultRequest) GetReturnFiveCategorySpoofResult() *string {
	return s.ReturnFiveCategorySpoofResult
}

func (s *CheckResultRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *CheckResultRequest) SetExtraImageControlList(v string) *CheckResultRequest {
	s.ExtraImageControlList = &v
	return s
}

func (s *CheckResultRequest) SetIsReturnImage(v string) *CheckResultRequest {
	s.IsReturnImage = &v
	return s
}

func (s *CheckResultRequest) SetMerchantBizId(v string) *CheckResultRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CheckResultRequest) SetReturnFiveCategorySpoofResult(v string) *CheckResultRequest {
	s.ReturnFiveCategorySpoofResult = &v
	return s
}

func (s *CheckResultRequest) SetTransactionId(v string) *CheckResultRequest {
	s.TransactionId = &v
	return s
}

func (s *CheckResultRequest) Validate() error {
	return dara.Validate(s)
}
