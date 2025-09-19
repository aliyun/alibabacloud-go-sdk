// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuaraFileDownloadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeQuaraFileDownloadInfoRequest
	GetFrom() *string
	SetQuaraFileId(v int32) *DescribeQuaraFileDownloadInfoRequest
	GetQuaraFileId() *int32
}

type DescribeQuaraFileDownloadInfoRequest struct {
	// The ID of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the quarantined file.
	//
	// > If you do not specify this parameter, you cannot call the RollbackSuspEventQuaraFile operation to restore a quarantined file. You can call the [DescribeSuspEventQuaraFiles](~~DescribeSuspEventQuaraFiles~~) operation to query the IDs of quarantined files.
	//
	// example:
	//
	// 123
	QuaraFileId *int32 `json:"QuaraFileId,omitempty" xml:"QuaraFileId,omitempty"`
}

func (s DescribeQuaraFileDownloadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuaraFileDownloadInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeQuaraFileDownloadInfoRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeQuaraFileDownloadInfoRequest) GetQuaraFileId() *int32 {
	return s.QuaraFileId
}

func (s *DescribeQuaraFileDownloadInfoRequest) SetFrom(v string) *DescribeQuaraFileDownloadInfoRequest {
	s.From = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoRequest) SetQuaraFileId(v int32) *DescribeQuaraFileDownloadInfoRequest {
	s.QuaraFileId = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoRequest) Validate() error {
	return dara.Validate(s)
}
