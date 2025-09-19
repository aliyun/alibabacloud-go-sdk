// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSuspEventQuaraFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *RollbackSuspEventQuaraFileRequest
	GetFrom() *string
	SetQuaraFileId(v int32) *RollbackSuspEventQuaraFileRequest
	GetQuaraFileId() *int32
	SetSourceIp(v string) *RollbackSuspEventQuaraFileRequest
	GetSourceIp() *string
}

type RollbackSuspEventQuaraFileRequest struct {
	// The ID of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the quarantined file.
	//
	// > If you do not configure this parameter, you cannot call the RollbackSuspEventQuaraFile operation to restore a quarantined file. You can call the [DescribeSuspEventQuaraFiles](~~DescribeSuspEventQuaraFiles~~) operation to query the IDs of quarantined files.
	//
	// example:
	//
	// 3921797
	QuaraFileId *int32 `json:"QuaraFileId,omitempty" xml:"QuaraFileId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.3.4
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s RollbackSuspEventQuaraFileRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackSuspEventQuaraFileRequest) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileRequest) GetFrom() *string {
	return s.From
}

func (s *RollbackSuspEventQuaraFileRequest) GetQuaraFileId() *int32 {
	return s.QuaraFileId
}

func (s *RollbackSuspEventQuaraFileRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *RollbackSuspEventQuaraFileRequest) SetFrom(v string) *RollbackSuspEventQuaraFileRequest {
	s.From = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetQuaraFileId(v int32) *RollbackSuspEventQuaraFileRequest {
	s.QuaraFileId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetSourceIp(v string) *RollbackSuspEventQuaraFileRequest {
	s.SourceIp = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) Validate() error {
	return dara.Validate(s)
}
