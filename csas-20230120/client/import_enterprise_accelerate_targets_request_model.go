// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportEnterpriseAccelerateTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEapId(v string) *ImportEnterpriseAccelerateTargetsRequest
	GetEapId() *string
	SetFileUrl(v string) *ImportEnterpriseAccelerateTargetsRequest
	GetFileUrl() *string
}

type ImportEnterpriseAccelerateTargetsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eap-6457339b546c4cfb
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://filename.xlsx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportEnterpriseAccelerateTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportEnterpriseAccelerateTargetsRequest) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseAccelerateTargetsRequest) GetEapId() *string {
	return s.EapId
}

func (s *ImportEnterpriseAccelerateTargetsRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImportEnterpriseAccelerateTargetsRequest) SetEapId(v string) *ImportEnterpriseAccelerateTargetsRequest {
	s.EapId = &v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsRequest) SetFileUrl(v string) *ImportEnterpriseAccelerateTargetsRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsRequest) Validate() error {
	return dara.Validate(s)
}
