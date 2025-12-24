// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordActionOutputListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionUuid(v string) *DescribeSoarRecordActionOutputListRequest
	GetActionUuid() *string
	SetLang(v string) *DescribeSoarRecordActionOutputListRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeSoarRecordActionOutputListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarRecordActionOutputListRequest
	GetPageSize() *int32
}

type DescribeSoarRecordActionOutputListRequest struct {
	// The UUID of the component action.
	//
	// >  You can call the [DescribeSoarTaskAndActions](~~DescribeSoarTaskAndActions~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2202c90d-fa93-4726-bc32-xxxxxx
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. If you leave this parameter empty, 10 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSoarRecordActionOutputListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListRequest) GetActionUuid() *string {
	return s.ActionUuid
}

func (s *DescribeSoarRecordActionOutputListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSoarRecordActionOutputListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarRecordActionOutputListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarRecordActionOutputListRequest) SetActionUuid(v string) *DescribeSoarRecordActionOutputListRequest {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetLang(v string) *DescribeSoarRecordActionOutputListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetPageNumber(v int32) *DescribeSoarRecordActionOutputListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetPageSize(v int32) *DescribeSoarRecordActionOutputListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) Validate() error {
	return dara.Validate(s)
}
