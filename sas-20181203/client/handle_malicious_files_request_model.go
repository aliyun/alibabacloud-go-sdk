// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleMaliciousFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIdList(v []*int64) *HandleMaliciousFilesRequest
	GetFileIdList() []*int64
	SetOperation(v string) *HandleMaliciousFilesRequest
	GetOperation() *string
}

type HandleMaliciousFilesRequest struct {
	// The list of file IDs to process.
	//
	// > You can call [ListAgentlessMaliciousFiles](~~ListAgentlessMaliciousFiles~~) to obtain the IDs.
	//
	// > -.
	FileIdList []*int64 `json:"FileIdList,omitempty" xml:"FileIdList,omitempty" type:"Repeated"`
	// The type of operation. Valid values:
	//
	// - addWhitelist: adds to the whitelist.
	//
	// - offWhitelist: removes from the whitelist.
	//
	// example:
	//
	// addWhitelist
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s HandleMaliciousFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleMaliciousFilesRequest) GoString() string {
	return s.String()
}

func (s *HandleMaliciousFilesRequest) GetFileIdList() []*int64 {
	return s.FileIdList
}

func (s *HandleMaliciousFilesRequest) GetOperation() *string {
	return s.Operation
}

func (s *HandleMaliciousFilesRequest) SetFileIdList(v []*int64) *HandleMaliciousFilesRequest {
	s.FileIdList = v
	return s
}

func (s *HandleMaliciousFilesRequest) SetOperation(v string) *HandleMaliciousFilesRequest {
	s.Operation = &v
	return s
}

func (s *HandleMaliciousFilesRequest) Validate() error {
	return dara.Validate(s)
}
