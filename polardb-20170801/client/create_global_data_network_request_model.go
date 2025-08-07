// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDataNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGlobalDataNetworkRequest
	GetDescription() *string
	SetDestinationFileSystemPath(v string) *CreateGlobalDataNetworkRequest
	GetDestinationFileSystemPath() *string
	SetDestinationId(v string) *CreateGlobalDataNetworkRequest
	GetDestinationId() *string
	SetDestinationRegion(v string) *CreateGlobalDataNetworkRequest
	GetDestinationRegion() *string
	SetDestinationType(v string) *CreateGlobalDataNetworkRequest
	GetDestinationType() *string
	SetFreezeSourceDuringSync(v string) *CreateGlobalDataNetworkRequest
	GetFreezeSourceDuringSync() *string
	SetSourceFileSystemPath(v string) *CreateGlobalDataNetworkRequest
	GetSourceFileSystemPath() *string
	SetSourceId(v string) *CreateGlobalDataNetworkRequest
	GetSourceId() *string
	SetSourceRegion(v string) *CreateGlobalDataNetworkRequest
	GetSourceRegion() *string
	SetSourceType(v string) *CreateGlobalDataNetworkRequest
	GetSourceType() *string
}

type CreateGlobalDataNetworkRequest struct {
	// example:
	//
	// mygdn
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// /
	DestinationFileSystemPath *string `json:"DestinationFileSystemPath,omitempty" xml:"DestinationFileSystemPath,omitempty"`
	// example:
	//
	// pfs-xxx
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// example:
	//
	// cn-beijing
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// example:
	//
	// pfs
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// true
	FreezeSourceDuringSync *string `json:"FreezeSourceDuringSync,omitempty" xml:"FreezeSourceDuringSync,omitempty"`
	// example:
	//
	// /
	SourceFileSystemPath *string `json:"SourceFileSystemPath,omitempty" xml:"SourceFileSystemPath,omitempty"`
	// example:
	//
	// oss-xxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// cn-wulanchabu
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// example:
	//
	// oss
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateGlobalDataNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDataNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalDataNetworkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGlobalDataNetworkRequest) GetDestinationFileSystemPath() *string {
	return s.DestinationFileSystemPath
}

func (s *CreateGlobalDataNetworkRequest) GetDestinationId() *string {
	return s.DestinationId
}

func (s *CreateGlobalDataNetworkRequest) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *CreateGlobalDataNetworkRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *CreateGlobalDataNetworkRequest) GetFreezeSourceDuringSync() *string {
	return s.FreezeSourceDuringSync
}

func (s *CreateGlobalDataNetworkRequest) GetSourceFileSystemPath() *string {
	return s.SourceFileSystemPath
}

func (s *CreateGlobalDataNetworkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGlobalDataNetworkRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CreateGlobalDataNetworkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateGlobalDataNetworkRequest) SetDescription(v string) *CreateGlobalDataNetworkRequest {
	s.Description = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetDestinationFileSystemPath(v string) *CreateGlobalDataNetworkRequest {
	s.DestinationFileSystemPath = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetDestinationId(v string) *CreateGlobalDataNetworkRequest {
	s.DestinationId = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetDestinationRegion(v string) *CreateGlobalDataNetworkRequest {
	s.DestinationRegion = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetDestinationType(v string) *CreateGlobalDataNetworkRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetFreezeSourceDuringSync(v string) *CreateGlobalDataNetworkRequest {
	s.FreezeSourceDuringSync = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetSourceFileSystemPath(v string) *CreateGlobalDataNetworkRequest {
	s.SourceFileSystemPath = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetSourceId(v string) *CreateGlobalDataNetworkRequest {
	s.SourceId = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetSourceRegion(v string) *CreateGlobalDataNetworkRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) SetSourceType(v string) *CreateGlobalDataNetworkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateGlobalDataNetworkRequest) Validate() error {
	return dara.Validate(s)
}
