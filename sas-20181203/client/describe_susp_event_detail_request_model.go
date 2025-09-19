// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeSuspEventDetailRequest
	GetFrom() *string
	SetLang(v string) *DescribeSuspEventDetailRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v int64) *DescribeSuspEventDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *DescribeSuspEventDetailRequest
	GetSourceIp() *string
	SetSuspiciousEventId(v int32) *DescribeSuspEventDetailRequest
	GetSuspiciousEventId() *int32
}

type DescribeSuspEventDetailRequest struct {
	// The data source of the exception. Set the value to sas.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
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
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the ID.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 121.33.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the exception.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32750999
	SuspiciousEventId *int32 `json:"SuspiciousEventId,omitempty" xml:"SuspiciousEventId,omitempty"`
}

func (s DescribeSuspEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeSuspEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSuspEventDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeSuspEventDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventDetailRequest) GetSuspiciousEventId() *int32 {
	return s.SuspiciousEventId
}

func (s *DescribeSuspEventDetailRequest) SetFrom(v string) *DescribeSuspEventDetailRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetLang(v string) *DescribeSuspEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribeSuspEventDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetSourceIp(v string) *DescribeSuspEventDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetSuspiciousEventId(v int32) *DescribeSuspEventDetailRequest {
	s.SuspiciousEventId = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
