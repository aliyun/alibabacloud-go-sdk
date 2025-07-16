// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetContentJobIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetSheetContentJobIdResponseBody
	GetJobId() *string
	SetRequestId(v string) *GetSheetContentJobIdResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetSheetContentJobIdResponseBody
	GetStatus() *string
	SetVendorRequestId(v string) *GetSheetContentJobIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetSheetContentJobIdResponseBody
	GetVendorType() *string
}

type GetSheetContentJobIdResponseBody struct {
	// example:
	//
	// 14640056080
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// init
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetSheetContentJobIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetSheetContentJobIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSheetContentJobIdResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSheetContentJobIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetSheetContentJobIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetSheetContentJobIdResponseBody) SetJobId(v string) *GetSheetContentJobIdResponseBody {
	s.JobId = &v
	return s
}

func (s *GetSheetContentJobIdResponseBody) SetRequestId(v string) *GetSheetContentJobIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSheetContentJobIdResponseBody) SetStatus(v string) *GetSheetContentJobIdResponseBody {
	s.Status = &v
	return s
}

func (s *GetSheetContentJobIdResponseBody) SetVendorRequestId(v string) *GetSheetContentJobIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetSheetContentJobIdResponseBody) SetVendorType(v string) *GetSheetContentJobIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetSheetContentJobIdResponseBody) Validate() error {
	return dara.Validate(s)
}
