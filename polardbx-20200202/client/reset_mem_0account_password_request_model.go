// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetMem0AccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ResetMem0AccountPasswordRequest
	GetDBInstanceName() *string
	SetMem0ApiKey(v string) *ResetMem0AccountPasswordRequest
	GetMem0ApiKey() *string
	SetRegionId(v string) *ResetMem0AccountPasswordRequest
	GetRegionId() *string
}

type ResetMem0AccountPasswordRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The API key.
	//
	// example:
	//
	// aafdf2e7d0988ef623***
	Mem0ApiKey *string `json:"Mem0ApiKey,omitempty" xml:"Mem0ApiKey,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetMem0AccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetMem0AccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetMem0AccountPasswordRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ResetMem0AccountPasswordRequest) GetMem0ApiKey() *string {
	return s.Mem0ApiKey
}

func (s *ResetMem0AccountPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetMem0AccountPasswordRequest) SetDBInstanceName(v string) *ResetMem0AccountPasswordRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ResetMem0AccountPasswordRequest) SetMem0ApiKey(v string) *ResetMem0AccountPasswordRequest {
	s.Mem0ApiKey = &v
	return s
}

func (s *ResetMem0AccountPasswordRequest) SetRegionId(v string) *ResetMem0AccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetMem0AccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
