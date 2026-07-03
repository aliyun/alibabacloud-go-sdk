// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserSourceLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddUserSourceLogConfigResponseBodyData) *AddUserSourceLogConfigResponseBody
	GetData() *AddUserSourceLogConfigResponseBodyData
	SetRequestId(v string) *AddUserSourceLogConfigResponseBody
	GetRequestId() *string
}

type AddUserSourceLogConfigResponseBody struct {
	// The return value of the request.
	Data *AddUserSourceLogConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserSourceLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserSourceLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigResponseBody) GetData() *AddUserSourceLogConfigResponseBodyData {
	return s.Data
}

func (s *AddUserSourceLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserSourceLogConfigResponseBody) SetData(v *AddUserSourceLogConfigResponseBodyData) *AddUserSourceLogConfigResponseBody {
	s.Data = v
	return s
}

func (s *AddUserSourceLogConfigResponseBody) SetRequestId(v string) *AddUserSourceLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddUserSourceLogConfigResponseBodyData struct {
	// The detailed information about the SLS log.
	//
	// example:
	//
	// cn-shanghai.siem-project.siem-logstore
	DiplayLine *string `json:"DiplayLine,omitempty" xml:"DiplayLine,omitempty"`
	// Indicates whether the details of the log collection task are returned. Valid values:
	//
	// example:
	//
	// 0
	Displayed *bool `json:"Displayed,omitempty" xml:"Displayed,omitempty"`
	// Indicates whether the log is collected. Valid values:
	//
	// example:
	//
	// 0
	Imported *bool `json:"Imported,omitempty" xml:"Imported,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase Threat Analysis.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	SourceLogCode *string `json:"SourceLogCode,omitempty" xml:"SourceLogCode,omitempty"`
	// The code of the product.
	//
	// example:
	//
	// sas
	SourceProdCode *string `json:"SourceProdCode,omitempty" xml:"SourceProdCode,omitempty"`
	// The ID of the Alibaba Cloud account for which the logs are collected.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The name of the Alibaba Cloud account for which the logs are collected.
	//
	// example:
	//
	// sas_account_xxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s AddUserSourceLogConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddUserSourceLogConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigResponseBodyData) GetDiplayLine() *string {
	return s.DiplayLine
}

func (s *AddUserSourceLogConfigResponseBodyData) GetDisplayed() *bool {
	return s.Displayed
}

func (s *AddUserSourceLogConfigResponseBodyData) GetImported() *bool {
	return s.Imported
}

func (s *AddUserSourceLogConfigResponseBodyData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *AddUserSourceLogConfigResponseBodyData) GetSourceLogCode() *string {
	return s.SourceLogCode
}

func (s *AddUserSourceLogConfigResponseBodyData) GetSourceProdCode() *string {
	return s.SourceProdCode
}

func (s *AddUserSourceLogConfigResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *AddUserSourceLogConfigResponseBodyData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *AddUserSourceLogConfigResponseBodyData) SetDiplayLine(v string) *AddUserSourceLogConfigResponseBodyData {
	s.DiplayLine = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetDisplayed(v bool) *AddUserSourceLogConfigResponseBodyData {
	s.Displayed = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetImported(v bool) *AddUserSourceLogConfigResponseBodyData {
	s.Imported = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetMainUserId(v int64) *AddUserSourceLogConfigResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSourceLogCode(v string) *AddUserSourceLogConfigResponseBodyData {
	s.SourceLogCode = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSourceProdCode(v string) *AddUserSourceLogConfigResponseBodyData {
	s.SourceProdCode = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSubUserId(v int64) *AddUserSourceLogConfigResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSubUserName(v string) *AddUserSourceLogConfigResponseBodyData {
	s.SubUserName = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
