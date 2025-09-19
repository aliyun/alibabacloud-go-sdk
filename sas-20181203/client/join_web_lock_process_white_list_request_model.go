// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinWebLockProcessWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessPaths(v []*string) *JoinWebLockProcessWhiteListRequest
	GetProcessPaths() []*string
	SetUuids(v string) *JoinWebLockProcessWhiteListRequest
	GetUuids() *string
}

type JoinWebLockProcessWhiteListRequest struct {
	// The paths of the processes.
	ProcessPaths []*string `json:"ProcessPaths,omitempty" xml:"ProcessPaths,omitempty" type:"Repeated"`
	// The UUIDs of the servers on which the processes run. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// 0c1714dc-f7a3-4265-8364-7aa3fce8****,1cc45e7d-7698-4b2c-89d8-e8cba407****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s JoinWebLockProcessWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinWebLockProcessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *JoinWebLockProcessWhiteListRequest) GetProcessPaths() []*string {
	return s.ProcessPaths
}

func (s *JoinWebLockProcessWhiteListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *JoinWebLockProcessWhiteListRequest) SetProcessPaths(v []*string) *JoinWebLockProcessWhiteListRequest {
	s.ProcessPaths = v
	return s
}

func (s *JoinWebLockProcessWhiteListRequest) SetUuids(v string) *JoinWebLockProcessWhiteListRequest {
	s.Uuids = &v
	return s
}

func (s *JoinWebLockProcessWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
