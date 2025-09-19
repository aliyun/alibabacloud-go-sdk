// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeLogMetaRequest
	GetFrom() *string
	SetLang(v string) *DescribeLogMetaRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v int64) *DescribeLogMetaRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *DescribeLogMetaRequest
	GetSourceIp() *string
}

type DescribeLogMetaRequest struct {
	// The ID of the request source. Default value: **aegis**. Valid values:
	//
	// 	- **aegis**: Server Guard
	//
	// 	- **sas**: Security Center
	//
	// >  If you use Server Guard, set the value to **aegis**. If you use Security Center, set the value to **sas**.
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
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 123.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLogMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogMetaRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeLogMetaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLogMetaRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeLogMetaRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeLogMetaRequest) SetFrom(v string) *DescribeLogMetaRequest {
	s.From = &v
	return s
}

func (s *DescribeLogMetaRequest) SetLang(v string) *DescribeLogMetaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogMetaRequest) SetResourceDirectoryAccountId(v int64) *DescribeLogMetaRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeLogMetaRequest) SetSourceIp(v string) *DescribeLogMetaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLogMetaRequest) Validate() error {
	return dara.Validate(s)
}
