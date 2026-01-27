// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceVectorSupportStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceVectorSupportStatusRequest
	GetDBInstanceId() *string
	SetStatus(v string) *ModifyDBInstanceVectorSupportStatusRequest
	GetStatus() *string
}

type ModifyDBInstanceVectorSupportStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-2vc2bn5c5b7g6****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Scheduled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDBInstanceVectorSupportStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceVectorSupportStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceVectorSupportStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceVectorSupportStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyDBInstanceVectorSupportStatusRequest) SetDBInstanceId(v string) *ModifyDBInstanceVectorSupportStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceVectorSupportStatusRequest) SetStatus(v string) *ModifyDBInstanceVectorSupportStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyDBInstanceVectorSupportStatusRequest) Validate() error {
	return dara.Validate(s)
}
