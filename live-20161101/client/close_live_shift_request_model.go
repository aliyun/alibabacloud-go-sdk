// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseLiveShiftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CloseLiveShiftRequest
	GetAppName() *string
	SetDomainName(v string) *CloseLiveShiftRequest
	GetDomainName() *string
	SetOwnerId(v int64) *CloseLiveShiftRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CloseLiveShiftRequest
	GetRegionId() *string
	SetStreamName(v string) *CloseLiveShiftRequest
	GetStreamName() *string
}

type CloseLiveShiftRequest struct {
	// The name of the application to which the live stream belongs. Wildcards (*) are supported, which indicates all AppName values under the domain name. You can view the AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. Wildcards (*) are supported, which indicates all StreamName values under the AppName. You can view the StreamName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s CloseLiveShiftRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseLiveShiftRequest) GoString() string {
	return s.String()
}

func (s *CloseLiveShiftRequest) GetAppName() *string {
	return s.AppName
}

func (s *CloseLiveShiftRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CloseLiveShiftRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloseLiveShiftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseLiveShiftRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *CloseLiveShiftRequest) SetAppName(v string) *CloseLiveShiftRequest {
	s.AppName = &v
	return s
}

func (s *CloseLiveShiftRequest) SetDomainName(v string) *CloseLiveShiftRequest {
	s.DomainName = &v
	return s
}

func (s *CloseLiveShiftRequest) SetOwnerId(v int64) *CloseLiveShiftRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseLiveShiftRequest) SetRegionId(v string) *CloseLiveShiftRequest {
	s.RegionId = &v
	return s
}

func (s *CloseLiveShiftRequest) SetStreamName(v string) *CloseLiveShiftRequest {
	s.StreamName = &v
	return s
}

func (s *CloseLiveShiftRequest) Validate() error {
	return dara.Validate(s)
}
