// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveEditingIndexFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetLiveEditingIndexFileRequest
	GetAppName() *string
	SetDomainName(v string) *GetLiveEditingIndexFileRequest
	GetDomainName() *string
	SetProjectId(v string) *GetLiveEditingIndexFileRequest
	GetProjectId() *string
	SetStreamName(v string) *GetLiveEditingIndexFileRequest
	GetStreamName() *string
}

type GetLiveEditingIndexFileRequest struct {
	// The application name of the live stream.
	//
	// example:
	//
	// testrecord
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name of the live stream.
	//
	// example:
	//
	// test.alivecdn.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the live stream editing project.
	//
	// example:
	//
	// *****cb6307a4edea614d8b3f3c*****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s GetLiveEditingIndexFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingIndexFileRequest) GoString() string {
	return s.String()
}

func (s *GetLiveEditingIndexFileRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetLiveEditingIndexFileRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *GetLiveEditingIndexFileRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetLiveEditingIndexFileRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *GetLiveEditingIndexFileRequest) SetAppName(v string) *GetLiveEditingIndexFileRequest {
	s.AppName = &v
	return s
}

func (s *GetLiveEditingIndexFileRequest) SetDomainName(v string) *GetLiveEditingIndexFileRequest {
	s.DomainName = &v
	return s
}

func (s *GetLiveEditingIndexFileRequest) SetProjectId(v string) *GetLiveEditingIndexFileRequest {
	s.ProjectId = &v
	return s
}

func (s *GetLiveEditingIndexFileRequest) SetStreamName(v string) *GetLiveEditingIndexFileRequest {
	s.StreamName = &v
	return s
}

func (s *GetLiveEditingIndexFileRequest) Validate() error {
	return dara.Validate(s)
}
