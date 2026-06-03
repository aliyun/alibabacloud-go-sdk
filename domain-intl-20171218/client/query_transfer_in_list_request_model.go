// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferInListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryTransferInListRequest
	GetDomainName() *string
	SetLang(v string) *QueryTransferInListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryTransferInListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryTransferInListRequest
	GetPageSize() *int32
	SetSimpleTransferInStatus(v string) *QueryTransferInListRequest
	GetSimpleTransferInStatus() *string
	SetSubmissionEndDate(v int64) *QueryTransferInListRequest
	GetSubmissionEndDate() *int64
	SetSubmissionStartDate(v int64) *QueryTransferInListRequest
	GetSubmissionStartDate() *int64
	SetUserClientIp(v string) *QueryTransferInListRequest
	GetUserClientIp() *string
}

type QueryTransferInListRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SimpleTransferInStatus *string `json:"SimpleTransferInStatus,omitempty" xml:"SimpleTransferInStatus,omitempty"`
	SubmissionEndDate      *int64  `json:"SubmissionEndDate,omitempty" xml:"SubmissionEndDate,omitempty"`
	SubmissionStartDate    *int64  `json:"SubmissionStartDate,omitempty" xml:"SubmissionStartDate,omitempty"`
	UserClientIp           *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTransferInListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInListRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferInListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTransferInListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTransferInListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryTransferInListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTransferInListRequest) GetSimpleTransferInStatus() *string {
	return s.SimpleTransferInStatus
}

func (s *QueryTransferInListRequest) GetSubmissionEndDate() *int64 {
	return s.SubmissionEndDate
}

func (s *QueryTransferInListRequest) GetSubmissionStartDate() *int64 {
	return s.SubmissionStartDate
}

func (s *QueryTransferInListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTransferInListRequest) SetDomainName(v string) *QueryTransferInListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTransferInListRequest) SetLang(v string) *QueryTransferInListRequest {
	s.Lang = &v
	return s
}

func (s *QueryTransferInListRequest) SetPageNum(v int32) *QueryTransferInListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTransferInListRequest) SetPageSize(v int32) *QueryTransferInListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTransferInListRequest) SetSimpleTransferInStatus(v string) *QueryTransferInListRequest {
	s.SimpleTransferInStatus = &v
	return s
}

func (s *QueryTransferInListRequest) SetSubmissionEndDate(v int64) *QueryTransferInListRequest {
	s.SubmissionEndDate = &v
	return s
}

func (s *QueryTransferInListRequest) SetSubmissionStartDate(v int64) *QueryTransferInListRequest {
	s.SubmissionStartDate = &v
	return s
}

func (s *QueryTransferInListRequest) SetUserClientIp(v string) *QueryTransferInListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTransferInListRequest) Validate() error {
	return dara.Validate(s)
}
