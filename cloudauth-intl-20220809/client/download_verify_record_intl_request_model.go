// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVerifyRecordIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DownloadVerifyRecordIntlRequest
	GetBizType() *string
	SetCode(v string) *DownloadVerifyRecordIntlRequest
	GetCode() *string
	SetDownloadMode(v string) *DownloadVerifyRecordIntlRequest
	GetDownloadMode() *string
	SetParam(v string) *DownloadVerifyRecordIntlRequest
	GetParam() *string
	SetProductType(v string) *DownloadVerifyRecordIntlRequest
	GetProductType() *string
}

type DownloadVerifyRecordIntlRequest struct {
	// Business type:
	//
	// - INVOKE_STATISTICS
	//
	// - INVOKE_RECORD
	//
	// example:
	//
	// INVOKE_RECORD
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Query code.
	//
	// example:
	//
	// vrf_intl_verify_record_real_id_idv_invoke_statistics_query
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Download mode:
	//
	// - **async**: Asynchronous
	//
	// - **sync**: Synchronous
	//
	// example:
	//
	// async
	DownloadMode *string `json:"DownloadMode,omitempty" xml:"DownloadMode,omitempty"`
	// Parameters related to the export and download query task.
	//
	// example:
	//
	// {\\"productCode\\":\\"FACE_LIVENESS\\",\\"startDs\\":\\"20251121\\",\\"endDs\\":\\"20251128\\",\\"language\\":\\"en\\"}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// Product Code.
	//
	// example:
	//
	// KYC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DownloadVerifyRecordIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadVerifyRecordIntlRequest) GoString() string {
	return s.String()
}

func (s *DownloadVerifyRecordIntlRequest) GetBizType() *string {
	return s.BizType
}

func (s *DownloadVerifyRecordIntlRequest) GetCode() *string {
	return s.Code
}

func (s *DownloadVerifyRecordIntlRequest) GetDownloadMode() *string {
	return s.DownloadMode
}

func (s *DownloadVerifyRecordIntlRequest) GetParam() *string {
	return s.Param
}

func (s *DownloadVerifyRecordIntlRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DownloadVerifyRecordIntlRequest) SetBizType(v string) *DownloadVerifyRecordIntlRequest {
	s.BizType = &v
	return s
}

func (s *DownloadVerifyRecordIntlRequest) SetCode(v string) *DownloadVerifyRecordIntlRequest {
	s.Code = &v
	return s
}

func (s *DownloadVerifyRecordIntlRequest) SetDownloadMode(v string) *DownloadVerifyRecordIntlRequest {
	s.DownloadMode = &v
	return s
}

func (s *DownloadVerifyRecordIntlRequest) SetParam(v string) *DownloadVerifyRecordIntlRequest {
	s.Param = &v
	return s
}

func (s *DownloadVerifyRecordIntlRequest) SetProductType(v string) *DownloadVerifyRecordIntlRequest {
	s.ProductType = &v
	return s
}

func (s *DownloadVerifyRecordIntlRequest) Validate() error {
	return dara.Validate(s)
}
