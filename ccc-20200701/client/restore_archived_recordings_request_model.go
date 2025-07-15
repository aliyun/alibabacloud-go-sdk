// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreArchivedRecordingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v string) *RestoreArchivedRecordingsRequest
	GetContactIds() *string
	SetInstanceId(v string) *RestoreArchivedRecordingsRequest
	GetInstanceId() *string
}

type RestoreArchivedRecordingsRequest struct {
	// example:
	//
	// [\\"job-216750038017142784\\"]
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestoreArchivedRecordingsRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreArchivedRecordingsRequest) GoString() string {
	return s.String()
}

func (s *RestoreArchivedRecordingsRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *RestoreArchivedRecordingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestoreArchivedRecordingsRequest) SetContactIds(v string) *RestoreArchivedRecordingsRequest {
	s.ContactIds = &v
	return s
}

func (s *RestoreArchivedRecordingsRequest) SetInstanceId(v string) *RestoreArchivedRecordingsRequest {
	s.InstanceId = &v
	return s
}

func (s *RestoreArchivedRecordingsRequest) Validate() error {
	return dara.Validate(s)
}
