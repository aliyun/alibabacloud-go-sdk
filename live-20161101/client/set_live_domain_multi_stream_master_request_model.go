// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamMasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetLiveDomainMultiStreamMasterRequest
	GetAppName() *string
	SetDomain(v string) *SetLiveDomainMultiStreamMasterRequest
	GetDomain() *string
	SetOwnerId(v int64) *SetLiveDomainMultiStreamMasterRequest
	GetOwnerId() *int64
	SetStreamName(v string) *SetLiveDomainMultiStreamMasterRequest
	GetStreamName() *string
	SetUpstreamSequence(v string) *SetLiveDomainMultiStreamMasterRequest
	GetUpstreamSequence() *string
}

type SetLiveDomainMultiStreamMasterRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The unique identifier of the stream ingest.
	//
	// This parameter is required.
	//
	// example:
	//
	// teststream_***
	UpstreamSequence *string `json:"UpstreamSequence,omitempty" xml:"UpstreamSequence,omitempty"`
}

func (s SetLiveDomainMultiStreamMasterRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamMasterRequest) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamMasterRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetLiveDomainMultiStreamMasterRequest) GetDomain() *string {
	return s.Domain
}

func (s *SetLiveDomainMultiStreamMasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveDomainMultiStreamMasterRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *SetLiveDomainMultiStreamMasterRequest) GetUpstreamSequence() *string {
	return s.UpstreamSequence
}

func (s *SetLiveDomainMultiStreamMasterRequest) SetAppName(v string) *SetLiveDomainMultiStreamMasterRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterRequest) SetDomain(v string) *SetLiveDomainMultiStreamMasterRequest {
	s.Domain = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterRequest) SetOwnerId(v int64) *SetLiveDomainMultiStreamMasterRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterRequest) SetStreamName(v string) *SetLiveDomainMultiStreamMasterRequest {
	s.StreamName = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterRequest) SetUpstreamSequence(v string) *SetLiveDomainMultiStreamMasterRequest {
	s.UpstreamSequence = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterRequest) Validate() error {
	return dara.Validate(s)
}
