// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanDistDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CleanDistDataRequest
	GetAppId() *string
	SetDataName(v string) *CleanDistDataRequest
	GetDataName() *string
	SetDataVersion(v string) *CleanDistDataRequest
	GetDataVersion() *string
	SetEnsRegionId(v string) *CleanDistDataRequest
	GetEnsRegionId() *string
}

type CleanDistDataRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DataName    *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	DataVersion *string `json:"DataVersion,omitempty" xml:"DataVersion,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s CleanDistDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CleanDistDataRequest) GoString() string {
	return s.String()
}

func (s *CleanDistDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *CleanDistDataRequest) GetDataName() *string {
	return s.DataName
}

func (s *CleanDistDataRequest) GetDataVersion() *string {
	return s.DataVersion
}

func (s *CleanDistDataRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CleanDistDataRequest) SetAppId(v string) *CleanDistDataRequest {
	s.AppId = &v
	return s
}

func (s *CleanDistDataRequest) SetDataName(v string) *CleanDistDataRequest {
	s.DataName = &v
	return s
}

func (s *CleanDistDataRequest) SetDataVersion(v string) *CleanDistDataRequest {
	s.DataVersion = &v
	return s
}

func (s *CleanDistDataRequest) SetEnsRegionId(v string) *CleanDistDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CleanDistDataRequest) Validate() error {
	return dara.Validate(s)
}
